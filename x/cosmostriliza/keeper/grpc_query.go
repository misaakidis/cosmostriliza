package keeper

import (
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
)

var _ types.QueryServer = Keeper{}
