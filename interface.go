package gowrt

import (
	"encoding/json"
	"fmt"
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

func (c *Client) GetInterfaceConfiguration(id, name string) (InterfaceConfiguration, error) {
	params := map[string]interface{}{
		"config": "network",
		"type":   "interface",
	}
	call := NewRpcCall(id, "call", "uci", "get", params)
	response, err := c.ApiCall(call)
	if err != nil {
		return InterfaceConfiguration{}, err
	}

	var interfaces map[string]map[string]InterfaceConfiguration
	if err := json.Unmarshal(response, &interfaces); err != nil {
		return InterfaceConfiguration{}, err
	}

	if values, ok := interfaces["values"]; ok {
		if iface, ifok := values[name]; ifok {
			return iface, nil
		}
	}
	return InterfaceConfiguration{}, fmt.Errorf("interface %s not found", name)
}
