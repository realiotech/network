package types

import (
	"strings"
)

const (
	// MainnetChainID defines the RealioNetwork EIP155 chain ID for mainnet
	MainnetChainID = "realionetwork_3301"
	// TestnetChainID defines the RealioNetwork EIP155 chain ID for testnet
	TestnetChainID = "realionetwork_3300"
	// BaseDenom defines the base denomination
	BaseDenom = "ario"
)

// IsMainnet returns true if the chain-id has the RealioNetwork mainnet EIP155 chain prefix.
func IsMainnet(chainID string) bool {
	return strings.HasPrefix(chainID, MainnetChainID)
}

// IsTestnet returns true if the chain-id has the RealioNetwork testnet EIP155 chain prefix.
func IsTestnet(chainID string) bool {
	return strings.HasPrefix(chainID, TestnetChainID)
}
