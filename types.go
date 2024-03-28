package gowrt

import (
	"fmt"
	"strings"
)

// StringBool is a boolean type thatis represented by the strings '0' and '1'
// in the OpenWrt Ubus API
type StringBool bool

func (b StringBool) UnmarshalJSON(data []byte) error {
	strData := strings.ReplaceAll(string(data), "\"", "")
	switch strData {
	case "0":
		b = false
	case "1":
		b = true
	default:
		return fmt.Errorf("invalid boolean value: %s", strData)
	}
	return nil
}

func (b StringBool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte("\"1\""), nil
	}
	return []byte("\"0\""), nil
}
