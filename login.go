package gowrt

import (
	"encoding/json"
	"fmt"
)

func (c *Client) Login(username, password string) error {
	body := map[string]interface{}{"username": username, "password": password}
	call := NewRpcCall("1", "call", "session", "login", body)
	response, err := c.ApiCall(call)
	if err != nil {
		return fmt.Errorf("login api response: %s", err)
	}
	var ubusSession UbusSession
	err = json.Unmarshal(response, &ubusSession)
	if err != nil {
		return fmt.Errorf("unmarshal api response: %s", err)
	}
	c.ubusSession = ubusSession
	return nil
}
