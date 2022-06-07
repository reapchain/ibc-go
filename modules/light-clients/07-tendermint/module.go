package tendermint

import (
	"github.com/reapchain/ibc-go/v3/modules/light-clients/07-tendermint/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
