package keeper

import (
	"github.com/spore-engineering/CondorChain/x/condorchain/types"
)

var _ types.QueryServer = Keeper{}
