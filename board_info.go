package gowrt

import (
	"encoding/json"
	"fmt"
)

type BoardInfo struct {
	Model struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"model"`
	Led     interface{} `json:"led"`
	Network struct {
		Lan struct {
			Ports    []string `json:"ports"`
			Protocol string   `json:"protocol"`
		} `json:"lan"`
		Wan struct {
			Device   string `json:"device"`
			Protocol string `json:"protocol"`
		} `json:"wan"`
	} `json:"network"`
	System struct {
		CompatVersion string `json:"compat_version"`
	} `json:"system"`
}

func (c *client) GetBoardInfo(id string) (BoardInfo, error) {
	var boardInfo BoardInfo
	params := map[string]interface{}{"path": "/etc/board.json"}
	call := NewRpcCall(id, "call", "file", "read", params)
	response, err := c.ApiCall(call)
	if err != nil {
		return boardInfo, fmt.Errorf("get board info: %s", err)
	}
	var jsonResponse struct {
		Data string `json:"data"`
	}
	err = json.Unmarshal(response, &jsonResponse)
	if err != nil {
		return boardInfo, fmt.Errorf("decode board info response: %s", err)
	}
	err = json.Unmarshal([]byte(jsonResponse.Data), &boardInfo)
	if err != nil {
		return boardInfo, fmt.Errorf("decode board info: %s", err)
	}
	return boardInfo, nil
}
