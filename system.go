package gowrt

import (
	"encoding/json"
	"fmt"
)

type NtpConfiguration struct {
	Index           int        `json:"index"`
	Enabled         StringBool `json:"enabled"`
	EnableServer    StringBool `json:"enable_server"`
	ServerInterface string     `json:"interface"`
	Servers         []string   `json:"server"`
}

func (c *Client) GetNtpConfiguration() (NtpConfiguration, error) {
	var ntpConfiguration NtpConfiguration
	params := map[string]interface{}{"config": "system", "type": "timeserver"}
	call := NewRpcCall("call", "uci", "get", params)
	response, err := c.ApiCall(call)
	if err != nil {
		return ntpConfiguration, fmt.Errorf("get ntp configuration: %s", err)
	}
	var jsonResponse struct {
		Values map[string]NtpConfiguration `json:"values"`
	}
	err = json.Unmarshal(response, &jsonResponse)
	if err != nil {
		return ntpConfiguration, fmt.Errorf("decode ntp configuration response: %s", err)
	}
	return ntpConfiguration, nil
}

func (c *Client) SetNtpConfiguration(cfg NtpConfiguration) error {
	params := map[string]interface{}{
		"config":  "system",
		"section": "ntp",
		"values": map[string]interface{}{
			"enabled":       cfg.Enabled,
			"enable_server": cfg.EnableServer,
			"interface":     cfg.ServerInterface,
			"server":        cfg.Servers,
		},
	}
	call := NewRpcCall("call", "uci", "set", params)
	response, err := c.ApiCall(call)
	if err != nil {
		fmt.Printf("SET RESPONSE: %#v\n", response)
		return fmt.Errorf("set ntp configuration: %s", err)
	}

	params = map[string]interface{}{
		"config": "system",
	}
	commitCall := NewRpcCall("call", "uci", "commit", params)
	fmt.Printf("%#v\n", commitCall)
	response, err = c.ApiCall(commitCall)
	if err != nil {
		fmt.Printf("COMMIT RESPONSE: %#v\n", response)
		return fmt.Errorf("commit ntp configuration: %s", err)
	}
	_ = response
	return nil
}
