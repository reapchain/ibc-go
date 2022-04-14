package simulation

import (
	"math/rand"

	simtypes "github.com/reapchain/cosmos-sdk/types/simulation"

	"github.com/reapchain/ibc-go/modules/core/02-client/types"
)

// GenClientGenesis returns the default client genesis state.
func GenClientGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
