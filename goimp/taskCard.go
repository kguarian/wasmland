package main

import (
	"errors"
	"fmt"
)

var TASK_CARDS map[string][]*taskCard = make(map[string][]*taskCard)

const TASK_CARD_TEMPLATE = `
<div id="%s" class="%s">
        <textarea id="%s" rows="%d" cols="%d"></textarea>
        <div>
            <button id="%s" onclick="%s()" disabled="%t">%s</button>

            <button id="%s" onclick="%s()" disabled="%t">%s</button>
        </div>
    </div>
`

//TODO: Choose the one that works better.
func buildTaskCardTemplate(id, startID, stopID string, rows, cols int, startFunc, endFunc func()) (retstring string) {
	return
}

type taskCard struct {
	id   string
	name string
	info string
}

func createInnerHTML(t *taskCard) (innerHTML string, err error) {
	if t == nil {
		LAST_ERROR = "taskCard parameter was nil"
		err = errors.New(LAST_ERROR)
	}
	innerHTML = fmt.Sprintf("<%s>", func() {
		// var v []*taskCard
		// for _, v = range TASK_CARDS {

		// }
	})
	return
}
