package gowrt

type UbusStatusCode int64

const (
	UBUS_STATUS_OK                UbusStatusCode = 0
	UBUS_STATUS_INVALID_COMMAND   UbusStatusCode = 1
	UBUS_STATUS_INVALID_ARGUMENT  UbusStatusCode = 2
	UBUS_STATUS_METHOD_NOT_FOUND  UbusStatusCode = 3
	UBUS_STATUS_NOT_FOUND         UbusStatusCode = 4
	UBUS_STATUS_NO_DATA           UbusStatusCode = 5
	UBUS_STATUS_PERMISSION_DENIED UbusStatusCode = 6
	UBUS_STATUS_TIMEOUT           UbusStatusCode = 7
	UBUS_STATUS_NOT_SUPPORTED     UbusStatusCode = 8
	UBUS_STATUS_UNKNOWN_ERROR     UbusStatusCode = 9
	UBUS_STATUS_CONNECTION_FAILED UbusStatusCode = 10
	UBUS_STATUS_LAST              UbusStatusCode = 11
)

type UbusSession struct {
	UbusRpcSession string                 `json:"ubus_rpc_session"`
	Timeout        int                    `json:"timeout"`
	Expires        int                    `json:"expires"`
	Acls           map[string]interface{} `json:"acls"`
	Data           map[string]interface{} `json:"data"`
}
