package gowrt

type RpcStatusCode int64

const (
	RPC_STATUS_ERROR_PARSE              RpcStatusCode = -32700
	RPC_STATUS_ERROR_INVALID_REQUEST    RpcStatusCode = -32600
	RPC_STATUS_ERROR_INTERNAL           RpcStatusCode = -32603
	RPC_STATUS_ERROR_INVALID_PARAMETERS RpcStatusCode = -32602
	RPC_STATUS_ERROR_METHOD_NOT_FOUND   RpcStatusCode = -32601
	RPC_STATUS_ERROR_TIMEOUT            RpcStatusCode = -32003
	RPC_STATUS_ERROR_ACCESS_DENIED      RpcStatusCode = -32002
	RPC_STATUS_ERROR_SESSION_NOT_FOUND  RpcStatusCode = -32001
	RPC_STATUS_ERROR_OBJECT_NOT_FOUND   RpcStatusCode = -32000
)

type rpcCall struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type rpcResponse struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Result  []interface{} `json:"result"`
	Error   rpcError      `json:"error"`
}

type rpcError struct {
	Code    int
	Message string
}
