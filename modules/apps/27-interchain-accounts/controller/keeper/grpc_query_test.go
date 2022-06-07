package keeper_test

import (
	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/reapchain/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"
)

func (suite *KeeperTestSuite) TestQueryParams() {
	ctx := sdk.WrapSDKContext(suite.chainA.GetContext())
	expParams := types.DefaultParams()
	res, _ := suite.chainA.GetSimApp().ICAControllerKeeper.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().Equal(&expParams, res.Params)
}
