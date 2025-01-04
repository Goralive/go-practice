package pokeapi

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name          string
		timeout       time.Duration
		cacheInterval time.Duration
	}{
		{
			name:          "Create client with standard settings",
			timeout:       time.Second * 10,
			cacheInterval: time.Minute * 5,
		},
		{
			name:          "Create client with short timeout and cache interval",
			timeout:       time.Millisecond * 500,
			cacheInterval: time.Second * 30,
		},
		{
			name:          "Create client with zero timeout and cache interval",
			timeout:       0,
			cacheInterval: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.timeout, tt.cacheInterval)

			// Verify that the client was created with correct timeout
			if client.httpClient.Timeout != tt.timeout {
				t.Errorf("Expected timeout %v, got %v", tt.timeout, client.httpClient.Timeout)
			}
		})
	}
}
