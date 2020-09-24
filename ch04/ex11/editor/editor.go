package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// ReadInputWithEditor open editor and return text written by the editor
func ReadInputWithEditor(editor, initialText, filename string) (text string, err error) {
	if err = ioutil.WriteFile(filename, []byte(initialText), 0644); err != nil {
		return
	}
	defer os.Remove(filename)
	if err = openEditor(editor, filename); err != nil {
		return
	}
	byteText, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return string(byteText), nil

}

func openEditor(editor string, filename string) error {
	c := exec.Command(editor, filename)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
