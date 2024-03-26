package gowrt

import (
	"fmt"

	"github.com/google/uuid"
)

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
	Id     string    `json:"id"`
	Method string    `json:"method"`
	Params rpcParams `json:"params"`
}

type apiPayload struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type rpcParams struct {
	Path    string
	Method  string
	Message map[string]interface{}
}

type rpcResponse struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Result  []interface{} `json:"result"`
	Error   rpcError      `json:"error"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewRpcCall(method, path, ubusMethod string, body map[string]interface{}) rpcCall {
	return rpcCall{
		Id:     uuid.New().String(),
		Method: method,
		Params: rpcParams{
			Path:    path,
			Method:  ubusMethod,
			Message: body,
		},
	}
}

func (c rpcCall) validate() error {
	if c.Method != "call" {
		return fmt.Errorf("unsupported rpc method: %s", c.Method)
	}
	return nil
}

func (c rpcCall) toApiPayload(s UbusSession) apiPayload {
	return apiPayload{
		JsonRpc: "2.0",
		Id:      c.Id,
		Method:  c.Method,
		Params: []interface{}{
			s.UbusRpcSession,
			c.Params.Path,
			c.Params.Method,
			c.Params.Message,
		},
	}
}
