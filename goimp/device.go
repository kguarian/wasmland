package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/google/uuid"
)

//TODO: REVISE, consider character limit for device names and usernames.
var INVALIDUSERCHARS []rune = []rune{'\t', ' ', ':', '[', ']', '.'}

//TODO: COMPLETE THE STATEMENT BELOW AND IMPLEMENT
//const namelengthlimit int = [decide on a limit and put it here]

//for internal use, mainly. This is shared with the client as of 4/4/2021
type Device struct {
	Username    string    `json:"userid"`
	Devicename  string    `json:"devname"`
	Device_uuid uuid.UUID `json:"devid"`
	Ipaddr      net.IP    `json:"ipaddr"`
	Indata      chan byte `json:"-"`
	Outdata     chan byte `json:"-"`
	Online      bool      `json:"-"`
	DeviceType  int       `json:"dev_type"`
}

//Sent from client to main on initiated contact.
type DeviceInfo struct {
	Userid     string `json:"userid"`
	Devicename string `json:"devname"`
	DeviceType string `json:"dev_type"`
}

func (a *Device) Equal(b *Device) bool {
	return a.Username == b.Username && a.Devicename == b.Devicename && a.Device_uuid == b.Device_uuid && a.Ipaddr.Equal(b.Ipaddr) && a.Indata == b.Indata && a.Outdata == b.Outdata && a.Online == b.Online && a.DeviceType == b.DeviceType
}

//This is an incredibly useless wrapper function that harms no one.
func (d *Device) MarshalDevice() []byte {
	retinfo, err := json.Marshal(d)
	Errhandle_Log(err, ERRMSG_JSON_MARSHALL)
	if err != nil {
		return nil
	}
	return retinfo
}

//TODO: Actually implement a login protocol
func (d *Device) Login() {
	d.Online = true
}

//TODO: Implement a login/logout system.
func (d *Device) Logout() {
	d.Online = false
}

//TODO: Implement a networked form of this.
//RequestConsent sends a single-file consent-to-transfer request from one device to another
func (d *Device) RequestConsent(recipientdevice Device, c contentinfo) error {
	//pointer not nil; checked above
	if !recipientdevice.Online {
		SetConsoleColor(RED)
		fmt.Printf("device [%s] is offline. Request Canceled.\n", recipientdevice.Device_uuid)
		SetConsoleColor(RESET)
		return errors.New(ERRMSG_DEVICEOFFLINE)
	}
	return nil
}

//TODO: Networked form
func Approveconsent(authtoken byte, c *contentinfo) (string, error) {
	if authtoken == 'Y' {
		return c.Senderid.String() + ":" + c.Name + ":" + strconv.FormatInt(c.Sizebytes, 10), nil
	}
	return "error", errors.New(ERRMSG_CONSENT_TO_SEND)
}

//TODO: implement
func (d *Device) SendContent(c *contentinfo) {

}

func EvalName(name string) bool {
	for _, c := range name {
		for _, c2 := range INVALIDUSERCHARS {
			if c == c2 {
				return false
			}
		}
	}
	return true
}
