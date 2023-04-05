package simulation_test

import (
	"fmt"
	"testing"

	"github.com/reapchain/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	clienttypes "github.com/reapchain/ibc-go/v3/modules/core/02-client/types"
	connectiontypes "github.com/reapchain/ibc-go/v3/modules/core/03-connection/types"
	channeltypes "github.com/reapchain/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/reapchain/ibc-go/v3/modules/core/24-host"
	"github.com/reapchain/ibc-go/v3/modules/core/simulation"
	ibctmtypes "github.com/reapchain/ibc-go/v3/modules/light-clients/07-tendermint/types"
	"github.com/reapchain/ibc-go/v3/testing/simapp"
)

func TestDecodeStore(t *testing.T) {
	app := simapp.Setup(false)
	dec := simulation.NewDecodeStore(*app.IBCKeeper)

	clientID := "clientidone"
	connectionID := "connectionidone"
	channelID := "channelidone"
	portID := "portidone"

	clientState := &ibctmtypes.ClientState{
		FrozenHeight: clienttypes.NewHeight(0, 10),
	}
	connection := connectiontypes.ConnectionEnd{
		ClientId: "clientidone",
		Versions: []*connectiontypes.Version{connectiontypes.NewVersion("1", nil)},
	}
	channel := channeltypes.Channel{
		State:   channeltypes.OPEN,
		Version: "1.0",
	}

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   host.FullClientStateKey(clientID),
				Value: app.IBCKeeper.ClientKeeper.MustMarshalClientState(clientState),
			},
			{
				Key:   host.ConnectionKey(connectionID),
				Value: app.IBCKeeper.Codec().MustMarshal(&connection),
			},
			{
				Key:   host.ChannelKey(portID, channelID),
				Value: app.IBCKeeper.Codec().MustMarshal(&channel),
			},
			{
				Key:   []byte{0x99},
				Value: []byte{0x99},
			},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"ClientState", fmt.Sprintf("ClientState A: %v\nClientState B: %v", clientState, clientState)},
		{"ConnectionEnd", fmt.Sprintf("ConnectionEnd A: %v\nConnectionEnd B: %v", connection, connection)},
		{"Channel", fmt.Sprintf("Channel A: %v\nChannel B: %v", channel, channel)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			if i == len(tests)-1 {
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			} else {
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
