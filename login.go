package gowrt

import (
	"encoding/json"
	"fmt"
)

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *client) Login(username, password string) error {
	body := loginBody{Username: username, Password: password}
	call := rpcCall{
		JsonRpc: "2.0",
		Id:      "1",
		Method:  "call",
		Params:  []interface{}{"session", "login", body},
	}
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
