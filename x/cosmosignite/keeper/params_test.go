package keeper_test

import (
	"testing"

	testkeeper "cosmos-ignite/testutil/keeper"
	"cosmos-ignite/x/cosmosignite/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmosigniteKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
