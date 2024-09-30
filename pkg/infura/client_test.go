package infura

import (
	"testing"
)

func TestClient_Subscribe(t *testing.T) {
	type fields struct {
		conf Config
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "",
			fields: fields{
				conf: Config{
					ApiKey:  "d60ab3ccc0024867b1e2087043bf8a6a",
					Host:    "mainnet.infura.io",
					Version: "v3",
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &Client{
				conf: tt.fields.conf,
			}

			cli.Subscribe()
		})
	}
}
