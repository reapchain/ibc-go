package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	
	"github.com/reapchain/ibc-go/modules/core/exported"
	icatypes "github.com/reapchain/ibc-go/modules/apps/27-interchain-accounts/types"
)

// EmitWriteErrorAcknowledgementEvent emits an event signalling an error acknowledgement and including the error details
func EmitWriteErrorAcknowledgementEvent(ctx sdk.Context, packet exported.PacketI, err error) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			icatypes.EventTypePacket,
			sdk.NewAttribute(sdk.AttributeKeyModule, icatypes.ModuleName),
			sdk.NewAttribute(icatypes.AttributeKeyAckError, err.Error()),
			sdk.NewAttribute(icatypes.AttributeKeyHostChannelID, packet.GetDestChannel()),
		),
	)
}
