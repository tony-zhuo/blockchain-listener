package infura

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"sync"
)

var (
	client     *Client
	clientOnce sync.Once
)

type Config struct {
	ApiKey  string
	Host    string
	Version string
}

type Client struct {
	conf Config
}

func NewClient(conf Config) {
	clientOnce.Do(func() {
		client = &Client{
			conf: conf,
		}
	})
}

func GetClient() *Client {
	if client == nil {
		panic("infura client not init.")
	}
	return client
}

func (cli *Client) Subscribe() {
	// 目标 WebSocket 服务器的地址
	u := url.URL{Scheme: "wss", Host: cli.conf.Host, Path: cli.path()}
	// 建立 WebSocket 连接
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// 订阅新区块头
	subscribeMessage := `{
		"jsonrpc": "2.0",
		"method": "eth_subscribe",
		"params": ["newHeads"],
		"id": 1
	}`
	err = c.WriteMessage(websocket.TextMessage, []byte(subscribeMessage))
	if err != nil {
		log.Fatal("write:", err)
	}

	// 持续监听新区块
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
		}

		fmt.Printf("New block received: %s\n", message)

		//// 可以通过 `eth_getBlockByHash` 获取区块详情，提取交易
		//blockHash := extractBlockHashFromMessage(message)
		//getBlockByHashMessage := fmt.Sprintf(`{
		//	"jsonrpc": "2.0",
		//	"method": "eth_getBlockByHash",
		//	"params": ["%s", true],
		//	"id": 1
		//}`, blockHash)
		//
		//err = c.WriteMessage(websocket.TextMessage, []byte(getBlockByHashMessage))
		//if err != nil {
		//	log.Fatal("write:", err)
		//}
		//
		//// 读取区块详情，解析交易
		//_, blockMessage, err := c.ReadMessage()
		//if err != nil {
		//	log.Fatal("read:", err)
		//}
		//fmt.Printf("Block details with transactions: %s\n", blockMessage)
	}
}

// 解析 WebSocket 消息，提取区块哈希 (这是一个示例函数)
func extractBlockHashFromMessage(message []byte) string {
	// 你需要使用 JSON 解析库解析消息并提取区块哈希
	// 此处仅作演示，实际需要解析返回的 JSON
	return "0xYourBlockHash"
}

func (cli *Client) path() string {
	//	wss://mainnet.infura.io/ws/v3/YOUR-API-KEY
	return fmt.Sprintf("/ws/%s/%s", cli.conf.Version, cli.conf.ApiKey)
}
