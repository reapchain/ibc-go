package types

import (
	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
	"github.com/reapchain/ibc-go/modules/core/exported"
)

// RegisterInterfaces registers the reapchain concrete client-related
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
	registry.RegisterImplementations(
		(*exported.ConsensusState)(nil),
		&ConsensusState{},
	)
	registry.RegisterImplementations(
		(*exported.Header)(nil),
		&Header{},
	)
	registry.RegisterImplementations(
		(*exported.Misbehaviour)(nil),
		&Misbehaviour{},
	)
}
