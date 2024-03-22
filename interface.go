package gowrt

import (
	"encoding/json"
)

type InterfaceConfiguration struct {
	Index     int    `json:".index"`
	Name      string `json:".name"`
	Device    string `json:"device"`
	Proto     string `json:"proto"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Ipv6      string `json:"ipv6"`      // "0"
	Ip6Assign string `json:"ip6assign"` // 60 set only on "lan"
	// IpAddrs    []string `json:"ipaddr"`
	ReqAddress string `json:"reqaddress"`
	ReqPrefix  string `json:"reqprefix"`
}

func (c *client) GetInterfaceConfiguration(id, name string) (InterfaceConfiguration, error) {
	var iface InterfaceConfiguration
	params := map[string]interface{}{
		"config": "network",
		"type":   "interface",
	}
	call := NewRpcCall(id, "call", "uci", "get", params)
	response, err := c.ApiCall(call)
	if err != nil {
		return iface, err
	}

	var interfaces map[string]map[string]InterfaceConfiguration
	if err := json.Unmarshal(response, &interfaces); err != nil {
		return iface, err
	}
	for ifaceName, config := range interfaces["values"] {
		if ifaceName != name {
			continue
		}
		iface = config
	}
	return iface, nil
}
