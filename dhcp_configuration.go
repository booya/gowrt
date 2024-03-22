package gowrt

type DhcpConfiguration struct {
}

func (c *client) GetDhcpConfiguration() DhcpConfiguration {
	var DhcpConfiguration DhcpConfiguration
	call := rpcCall{
		JsonRpc: "2.0",
		Id:      "4",
		Method:  "call",
		Params:  []interface{}{"uci", "get", map[string]string{"config": "dhcp"}},
	}
	_ = call
	return DhcpConfiguration
}
