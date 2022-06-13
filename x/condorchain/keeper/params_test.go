package keeper_test

import (
	"testing"

	testkeeper "github.com/spore-engineering/CondorChain/testutil/keeper"
	"github.com/spore-engineering/CondorChain/x/condorchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CondorchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
