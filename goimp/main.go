package main

import (
	"fmt"
	"net"
	"syscall/js"
	"time"
)

var JS_GLOBAL js.Value = js.Global()
var display js.Value = JS_GLOBAL.Get("timer_display")

func main() {
	var lifetime chan JsFunction

	var resetTimerAPI_Hook, startTimerAPI_Hook, stopTimerAPI_Hook JsFunction

	var sendPingAPI_Hook JsFunction

	lifetime = make(chan JsFunction)

	startTimerAPI_Hook.Init("wasm_startTimer", startTimerAPI)
	stopTimerAPI_Hook.Init("wasm_stopTimer", stopTimerAPI)
	resetTimerAPI_Hook.Init("wasm_resetTimer", resetTimerAPI)
	sendPingAPI_Hook.Init("wasm_devconn_ping", sendPingAPI)

	startTimerAPI_Hook.Expose(JS_GLOBAL)
	stopTimerAPI_Hook.Expose(JS_GLOBAL)
	resetTimerAPI_Hook.Expose(JS_GLOBAL)
	sendPingAPI_Hook.Expose(JS_GLOBAL)

	display.Set("value", "12.15 - fixed stop")
	//TODO: The display is your instruction set
	_ = <-lifetime
}

/*
	usage: startTimerAPI(id js.Value from string)
*/
func startTimerAPI(this js.Value, val []js.Value) (retVal interface{}) {
	var id string
	var duration_asString string
	var refTime time.Time
	var inputDuration time.Duration
	var timr *timer
	var err error

	println("startTimerAPI start")
	fmt.Printf("val length: %d\n", len(val))
	if len(val) < 2 {
		println("insufficient arguments to startTimerAPI function. There should be 2.")
		retVal = js.ValueOf(1)
		return
	}
	id = val[0].String()

	if Const_timers[id] != nil && !Const_timers[id].stopped {
		println("no-op. timer already running.")
		return
	}
	//convert javascript string to a go-usable string
	duration_asString = val[1].String()
	//create timer
	refTime = time.Now()
	fmt.Printf("attempted string: %s\n", duration_asString)
	inputDuration, err = time.ParseDuration(duration_asString)
	if err != nil {
		println("couldn't parse user input for time duration")
		retVal = js.ValueOf(1)
		return
	}

	//"Everything with a space, everything in its space"
	timr = &timer{
		id:                 id,
		millisecondsPassed: 0,
		millisecondGoal:    inputDuration * time.Millisecond,
		startTime:          refTime,
		endTime:            refTime.Add(inputDuration),
		signals:            make(chan byte, 1),
	}

	Const_timers[timr.id] = timr

	println("calling startTimerGo")
	go runTimerGo(timr)
	println("startTimerAPI finished")
	return
}

//convert, parse, function call.
/*
usage: stopTimerAPI(id js.Value from string)
*/
func stopTimerAPI(this js.Value, val []js.Value) (retVal interface{}) {
	var timr *timer
	var id string
	var errorCodeChannel chan byte
	if len(val) == 0 {
		retVal = js.ValueOf(1)
		return
	}
	errorCodeChannel = make(chan byte)
	id = val[0].String()
	timr = Const_timers[id]
	if timr == nil {
		return js.ValueOf(1)
	}
	go stopTimerGo(timr, errorCodeChannel)
	println(id)
	return
}

/*
usage: resetTimerAPI(id js.Value from string)

RETURN VALUE USED HERE!!! DO NOT CHANGE BEFORE READING COMMENTS!!!
*/
func resetTimerAPI(this js.Value, val []js.Value) (retVal interface{}) {
	var id string

	if len(val) == 0 {
		return js.ValueOf(1)
	}

	id = val[0].String()
	if Const_timers[id] == nil {
		println("timer not found")
		//indicates that we should attempt to restart the timer
		//according to then text field
		return js.ValueOf(1)
	}
	resetTimerGo(Const_timers[id])
	//implies successful reset
	return js.ValueOf(0)
}

func sendPingAPI(this js.Value, val []js.Value) (retVal interface{}) {
	var devid, userid string
	var ip_addr net.IP
	if len(val) < 3 {
		println("empty argument list")
		return
	}
	ip_addr = net.ParseIP(val[2].String())
	if ip_addr == nil {
		println("invalid ip address")
		return
	}

	devid = val[0].String()
	userid = val[1].String()
	go sendPingGo(devid, userid, ip_addr)
	return
}
