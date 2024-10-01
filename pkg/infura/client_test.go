package infura

import (
	"testing"
)

func TestClient_Subscribe(t *testing.T) {
	NewClient(Config{
		ApiKey:  "d60ab3ccc0024867b1e2087043bf8a6a",
		Host:    "mainnet.infura.io",
		Version: "v3",
	})
	cli := GetClient()
	
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli.Subscribe()
		})
	}
}
