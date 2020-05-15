package map2value

import (
	"errors"
	"net"
)

func MapToNetIP(d map[string]interface{}, key string) (net.IP, error) {
	var v4fieldname, v6fieldname string
	switch key {
	case "src":
		v4fieldname = "sourceIPv4Address"
		v6fieldname = "sourceIPv6Address"
	case "dst":
		v4fieldname = "destinationIPv4Address"
		v6fieldname = "destinationIPv6Address"
	case "client":
		v4fieldname = "conn_client_ipv4_address"
		v6fieldname = "conn_client_ipv6_address"
	case "server":
		v4fieldname = "conn_server_ipv4_address"
		v6fieldname = "conn_server_ipv6_address"
	case "postnatsrc":
		v4fieldname = "postNATSourceIPv4Address"
		v6fieldname = "postNATSourceIPv6Address"
	case "postnatdst":
		v4fieldname = "postNATDestinationIPv4Address"
		v6fieldname = "postNATDestinationIPv6Address"
	}
	if v4, valid4 := d[v4fieldname]; valid4 {
		return v4.(net.IP), nil
	} else if v6, valid6 := d[v6fieldname]; valid6 {
		return v6.(net.IP), nil
	} else {
		return nil, errors.New("Key is not exist")
	}
}

func MapToNetPort(d map[string]interface{}, key string) uint16 {
	var fieldname string
	switch key {
	case "src":
		fieldname = "sourceTransportPort"
	case "dst":
		fieldname = "destinationTransportPort"
	case "client":
		fieldname = "conn_client_trans_port"
	case "server":
		fieldname = "conn_server_trans_port"
	case "postnatsrc":
		fieldname = "postNAPTSourceTransportPort"
	case "postnatdst":
		fieldname = "postNAPTDestinationTransportPort"
	}
	v, valid := d[fieldname]
	if valid {
		return v.(uint16)
	}
	return 0
}
