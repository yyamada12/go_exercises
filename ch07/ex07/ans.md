理由：デフォルト値が String()メソッドを通して表示されているため。

詳細：
flag.go のコードより、ヘルプメッセージの出力時にはデフォルト値を(flag の型が string でない場合)
`fmt.Sprintf(" (default %v)", flag.DefValue)` として出力している。

```
// PrintDefaults prints, to standard error unless configured otherwise, the
// default values of all defined command-line flags in the set. See the
// documentation for the global function PrintDefaults for more information.
func (f *FlagSet) PrintDefaults() {
	f.VisitAll(func(flag *Flag) {
		s := fmt.Sprintf("  -%s", flag.Name) // Two spaces before -; see next two comments.
		name, usage := UnquoteUsage(flag)
		if len(name) > 0 {
			s += " " + name
		}
		// Boolean flags of one ASCII letter are so common we
		// treat them specially, putting their usage on the same line.
		if len(s) <= 4 { // space, space, '-', 'x'.
			s += "\t"
		} else {
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			s += "\n    \t"
		}
		s += strings.ReplaceAll(usage, "\n", "\n    \t")

		if !isZeroValue(flag, flag.DefValue) {
			if _, ok := flag.Value.(*stringValue); ok {
				// put quotes on the value
				s += fmt.Sprintf(" (default %q)", flag.DefValue)
			} else {
				s += fmt.Sprintf(" (default %v)", flag.DefValue)
			}
		}
		fmt.Fprint(f.Output(), s, "\n")
	})
}
```

print.go のコードより、 %v verb による出力では型 switch を用いて
引数が Stringer インターフェースを満足している場合、String メソッドを用いて表示している。

```
// If a string is acceptable according to the format, see if
// the value satisfies one of the string-valued interfaces.
// Println etc. set verb to %v, which is "stringable".
switch verb {
case 'v', 's', 'x', 'X', 'q':
    // Is it an error or Stringer?
    // The duplication in the bodies is necessary:
    // setting handled and deferring catchPanic
    // must happen before calling the method.
    switch v := p.arg.(type) {
    case error:
        handled = true
        defer p.catchPanic(p.arg, verb, "Error")
        p.fmtString(v.Error(), verb)
        return

    case Stringer:
        handled = true
        defer p.catchPanic(p.arg, verb, "String")
        p.fmtString(v.String(), verb)
        return
    }
}
```
