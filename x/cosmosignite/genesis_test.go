package cosmosignite_test

import (
	"testing"

	keepertest "cosmos-ignite/testutil/keeper"
	"cosmos-ignite/testutil/nullify"
	"cosmos-ignite/x/cosmosignite"
	"cosmos-ignite/x/cosmosignite/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosigniteKeeper(t)
	cosmosignite.InitGenesis(ctx, *k, genesisState)
	got := cosmosignite.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
