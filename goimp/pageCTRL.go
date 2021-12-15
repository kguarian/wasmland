package main

import (
	"fmt"
	"syscall/js"
)

//Desired functionality: Ability to position elements relatively and absolutely, insert/add new divs and elements and such, edit innerHTML.
type Page struct {
	Title       string
	Description string
}

type Widget struct {
	Tag           string
	Id            string
	Type          string
	Functionality map[string]JsFunction
	Associations  map[string]string
	Subs          []*Widget
}

type JsFunction struct {
	Name     string
	Function func(this js.Value, val []js.Value) (i interface{})
}

var FUNC_TABLE map[string]func(this js.Value, val []js.Value) (i interface{}) = make(map[string]func(this js.Value, val []js.Value) (i interface{}))

func (j *JsFunction) Init(name string, f func(this js.Value, val []js.Value) (i interface{})) {
	j.Name = name
	j.Function = f
}

func (j *JsFunction) Expose(GlobalScope js.Value) {
	GlobalScope.Set(j.Name, js.FuncOf(j.Function))
	FUNC_TABLE[j.Name] = j.Function
}

func (w *Widget) CreateAssociations() (retString string) {
	// var k string
	// var v func()
	// for k, v = range w.Functionality {

	// }
	return retString
}
func (w *Widget) CreateFunctionality() (retString string) {
	return retString
}

func (w *Widget) Create() (retString string) {
	retString = fmt.Sprintf("<%s id=\"%s\" type=\"%s\" %s %s %s></%s>", w.Tag, w.Id, w.Type, w.CreateAssociations(), w.CreateFunctionality())
	return
}
