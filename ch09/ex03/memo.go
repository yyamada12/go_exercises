// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

// Func is the type of the function to memoize.
type Func func(key string, done <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res      result
	ready    chan struct{} // closed when res is ready
	handover chan callReqs // used when Get canceled
}

type callReqs struct {
	f     Func
	key   string
	cache map[string]*entry
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	done     <-chan struct{}
	response chan<- result // the client wants a single result
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, done, response}
	select {
	case res := <-response:
		return res.value, res.err
	case <-done:
		return nil, nil
	}
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{}), handover: make(chan callReqs)}
			cache[req.key] = e
			go e.call(f, req.key, req.done, cache) // call f(key)
		}
		go e.deliver(req.response, req.done)
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}, cache map[string]*entry) {
	ch := make(chan struct{})
	// Evaluate the function.
	go func() {
		e.res.value, e.res.err = f(key, done)
		ch <- struct{}{}
	}()
	// Broadcast the ready condition.
	select {
	case <-ch:
		close(e.ready)
	case <-done:
		select {
		case e.handover <- callReqs{f, key, cache}:
			// hand over calling Func to another Get
		default:
			delete(cache, key)
		}
	}
}

func (e *entry) deliver(response chan<- result, done <-chan struct{}) {
	select {
	// Wait for the ready condition.
	case <-e.ready:
		// Send the result to the client.
		response <- e.res
	case callreqs := <-e.handover:
		go e.call(callreqs.f, callreqs.key, done, callreqs.cache)
		go e.deliver(response, done)
	case <-done:
		//do nothing
	}
}
