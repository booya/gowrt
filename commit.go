package gowrt

import (
	"fmt"
)

func (c *Client) UciCommit(section string) error {
	params := map[string]interface{}{
		"config": section,
	}
	call := NewRpcCall("call", "uci", "commit", params)
	response, err := c.ApiCall(call)
	if err != nil {
		return fmt.Errorf("get board info: %s", err)
	}
	_ = response
	return nil
}
