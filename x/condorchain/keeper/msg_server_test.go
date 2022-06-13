package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/spore-engineering/CondorChain/testutil/keeper"
	"github.com/spore-engineering/CondorChain/x/condorchain/keeper"
	"github.com/spore-engineering/CondorChain/x/condorchain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CondorchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
