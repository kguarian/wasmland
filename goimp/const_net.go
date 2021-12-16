//Networking constants

package main

const (
	//only works at home
	//IP_main          = "192.168.1.118:8081"
	NETCODE_ERR      = byte(0)
	NETCODE_SUC      = byte(1)
	PORT_LOWER_BOUND = 49152
	PORT_UPPER_BOUND = 65535
	TCP              = "tcp"
)

const (
	nativegoclient int = iota
	jsclient
	cclient
)

const (
	NETMSG_TYPE_TEXT = iota
	NETMSG_TYPE_BINARY
	NETMSG_TYPE_UNSPECIFIED
)

const (
	PORT = ":8081"
	//Requests from device
	NETREQ_NEWDEVICE               = "nd"
	NETREQ_NEWDEVICE_JAVASCRIPT    = "nd_js"
	NETREQ_NEWCONTENTTRANSFER      = "nct"
	NETREQ_TRUSTREQUEST            = "tc"
	NETREQ_TRUSTCONSENT_JAVASCRIPT = "tc_js"
	NETMSG_SECCUSSFUL_TRANSMISSION = "gxd_trans"
	TIMEOUT_ERROR                  = "error_timeout"
)

const (
	//Directive from main
	NETDIR_SENDCONTENTINFO = "s_ci"
)

var (
	IP_main = "192.168.1.118:8081"
)
