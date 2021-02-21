package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GameAll(c context.Context, req *types.QueryAllGameRequest) (*types.QueryAllGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var games []*types.Game
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	gameStore := prefix.NewStore(store, types.KeyPrefix(types.GameKey))

	pageRes, err := query.Paginate(gameStore, req.Pagination, func(key []byte, value []byte) error {
		var game types.Game
		if err := k.cdc.UnmarshalBinaryBare(value, &game); err != nil {
			return err
		}

		games = append(games, &game)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGameResponse{Game: games, Pagination: pageRes}, nil
}

func (k Keeper) Game(c context.Context, req *types.QueryGetGameRequest) (*types.QueryGetGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var game types.Game
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.GameKey+req.Id)), &game)

	return &types.QueryGetGameResponse{Game: &game}, nil
}
