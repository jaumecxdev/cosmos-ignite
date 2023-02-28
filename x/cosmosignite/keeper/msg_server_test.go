package keeper_test

import (
	"context"
	"testing"

	keepertest "cosmos-ignite/testutil/keeper"
	"cosmos-ignite/x/cosmosignite/keeper"
	"cosmos-ignite/x/cosmosignite/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosigniteKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
