package main

import (
	"fmt"
	"syscall/js"
	"time"
)

var Const_timers map[string]*timer = make(map[string]*timer)

type timer struct {
	millisecondsPassed time.Duration
	millisecondGoal    time.Duration
	startTime          time.Time
	endTime            time.Time
	id                 string
	//event queue. Church-Turing Thesis: 1 channel can do anything that parallel channels can do, but possibly (and often) slower.
	//also, this is a reference type, treating like pointer
	signals chan byte
	stopped bool
	display *[]string
}

func runTimerGo(t *timer) (errcode int) {
	//reference type
	var synValue byte
	var duration time.Duration

	fmt.Printf("%v", *t)
	// var dur time.Duration

	fmt.Println(t.endTime)

	for ; time.Now().Before(t.endTime) || t.stopped; time.Sleep(time.Millisecond) {
		// dur = time.Until(t.endTime)
		if !t.stopped {
			JS_GLOBAL.Get("timer_display").Set("value", js.ValueOf(time.Until(t.endTime).String()))
		}
		// t.mut.Lock()
		// defer t.mut.Unlock()
		if len(t.signals) != 0 {
			synValue = <-t.signals
			switch synValue {
			//stop funcall triggered.
			case 0:
				t.stopped = true
				break
			case 1:
				duration = t.endTime.Sub(t.startTime)
				t.millisecondsPassed = 0
				t.millisecondGoal = duration
				t.startTime = time.Now()
				t.endTime = t.startTime.Add(t.millisecondGoal)
				JS_GLOBAL.Get("timer_display").Set("value", duration.String())
				//no need to exit the function, right?
				break
			default:
				println("WOAH! This should never happen! FIX IN TIMER.GO")
				break
			}
		}
	}

	JS_GLOBAL.Get("timer_display").Set("value", js.ValueOf("Timer Finished"))
	Const_timers[t.id] = nil
	return
}

func stopTimerGo(timr *timer, ecc chan byte) (errcode int) {
	println("stopTimerGo()")
	var callTime time.Time

	callTime = time.Now()

	if timr == nil {
		fmt.Println("timr == nil")
		errcode = 1
		return
	}

	if timr.endTime.Before(callTime) {

		fmt.Println("timr.endTime.Before(callTime)")
		errcode = 2
		return
	}

	//timer should be handled via the channel in this stop event.
	timr.signals <- 0
	fmt.Printf("passed stop token (0) into signal channel\n")
	return
}

func resetTimerGo(timr *timer) {
	println("ResetTimerGo()")
	if Const_timers[timr.id] == nil {
		println("timer doesn't exist")
	}
	timr.signals <- 1
	println("passed reset token (1) into signal channel")
}

func updateTimerGo(timr *timer) (errcode int) {
	println("updateTimerGo()")
	var duration_ms uint64

	JS_GLOBAL.Get(timr.id).Set("value", fmt.Sprintf("%f seconds", float64(duration_ms)/1000))
	return
}
