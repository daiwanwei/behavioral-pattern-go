package main

import (
	"behavioral-pattern-go/memento/emaileditor"
	"fmt"
)

func main() {
	editor := emaileditor.NewEditor()
	fmt.Println(editor.Current.GetState())
	editor.Current.SetTitle("hello")
	fmt.Println(editor.Current.GetState())
	editor.Save()
	editor.Current.SetTitle("hello2")
	fmt.Println(editor.Current.GetState())
	editor.Save()
	editor.Current.SetTitle("hello3")
	fmt.Println(editor.Current.GetState())
	editor.Undo()
	fmt.Println(editor.Current.GetState())
}
