package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"syscall/js"
)

func sendPingGo(devid, userid string, ip_addr net.IP) {
	var devinf DeviceInfo
	var received_device Device

	//var connServer net.Conn
	var err error
	var b []byte

	//because this function exits before network processes finish:
	println(ip_addr.String())
	ws := js.Global().Get("WebSocket").New("ws://" + ip_addr.String() + ":8081/")
	ws.Set("binaryType", "arraybuffer")
	ws.Call("addEventListener", "open", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		protocol := NewNetmessage(NETMSG_TYPE_TEXT, NETREQ_NEWDEVICE_JAVASCRIPT)
		b, err = json.Marshal(&protocol)
		if err != nil {
			log.Printf("%s\n", err)
			return nil
		}
		jsvalue := js.Global().Get("Uint8Array").New(len(b))
		js.CopyBytesToJS(jsvalue, b)
		ws.Call("send", jsvalue)

		devinf = DeviceInfo{Userid: userid, Devicename: devid}
		b, err = json.Marshal(&devinf)

		if err != nil {
			Errhandle_Log(err, ERRMSG_JSON_MARSHALL)
		}
		jsvalue = js.Global().Get("Uint8Array").New(len(b))
		js.CopyBytesToJS(jsvalue, b)
		ws.Call("send", jsvalue)
		return nil
	}))

	ws.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		jsvalue := js.Global().Get("Uint8Array").New(args[0].Get("data"))
		jslength := args[0].Get("data").Get("byteLength").Int()

		b = make([]byte, jslength)
		js.CopyBytesToGo(b, jsvalue)

		err = json.Unmarshal(b, &received_device)

		if err != nil {
			Errhandle_Log(err, ERRMSG_JSON_MARSHALL)
		}

		log.Printf("received device: %v", received_device)
		js.Global().Get("document").Call("getElementById", "deviceid").Set("innerHTML", js.ValueOf(fmt.Sprintf("%v", received_device)))
		return nil
	}))

	return
}
