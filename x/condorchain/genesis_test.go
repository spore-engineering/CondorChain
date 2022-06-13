package condorchain_test

import (
	"testing"

	keepertest "github.com/spore-engineering/CondorChain/testutil/keeper"
	"github.com/spore-engineering/CondorChain/testutil/nullify"
	"github.com/spore-engineering/CondorChain/x/condorchain"
	"github.com/spore-engineering/CondorChain/x/condorchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CondorchainKeeper(t)
	condorchain.InitGenesis(ctx, *k, genesisState)
	got := condorchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
