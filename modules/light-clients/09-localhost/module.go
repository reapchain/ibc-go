package localhost

import (
	"github.com/reapchain/ibc-go/v3/modules/light-clients/09-localhost/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
