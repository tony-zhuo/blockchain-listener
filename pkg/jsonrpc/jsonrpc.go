package jsonrpc

import (
	"blockchain-listener/pkg/libs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
)

func Post(url, method string, params []interface{}) (*Response, error) {
	requestID := libs.RandomInt(math.MaxInt)
	requestBody := Request{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		ID:      requestID,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON request: %v", err)
	}

	// 发送 POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// 解析 JSON-RPC 响应
	var rpcResponse Response
	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON response: %v", err)
	}

	if rpcResponse.ID != requestBody.ID {
		return nil, fmt.Errorf("error id JSON response: %v", rpcResponse.ID)
	}

	return &rpcResponse, nil
}
