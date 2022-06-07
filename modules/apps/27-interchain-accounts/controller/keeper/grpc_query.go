package keeper

import (
	"context"

	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/reapchain/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"
)

var _ types.QueryServer = Keeper{}

// Params implements the Query/Params gRPC method
func (q Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{
		Params: &params,
	}, nil
}
