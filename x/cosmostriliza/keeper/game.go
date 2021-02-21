package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
	"strconv"
)

// GetGameCount get the total number of game
func (k Keeper) GetGameCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameCountKey))
	byteKey := types.KeyPrefix(types.GameCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetGameCount set the total number of game
func (k Keeper) SetGameCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameCountKey))
	byteKey := types.KeyPrefix(types.GameCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateGame creates a game with a new id and update the count
func (k Keeper) CreateGame(ctx sdk.Context, msg types.MsgCreateGame) {
	// Create the game
	count := k.GetGameCount(ctx)
	var game = types.Game{
		Creator:    msg.Creator,
		Id:         strconv.FormatInt(count, 10),
		Rows:       msg.Rows,
		Cols:       msg.Cols,
		Strike:     msg.Strike,
		Reward:     msg.Reward,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
	key := types.KeyPrefix(types.GameKey + game.Id)
	value := k.cdc.MustMarshalBinaryBare(&game)
	store.Set(key, value)

	// Update game count
	k.SetGameCount(ctx, count+1)
}

// SetGame set a specific game in the store
func (k Keeper) SetGame(ctx sdk.Context, game types.Game) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
	b := k.cdc.MustMarshalBinaryBare(&game)
	store.Set(types.KeyPrefix(types.GameKey+game.Id), b)
}

// GetGame returns a game from its id
func (k Keeper) GetGame(ctx sdk.Context, key string) types.Game {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
	var game types.Game
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.GameKey+key)), &game)
	return game
}

// HasGame checks if the game exists
func (k Keeper) HasGame(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
	return store.Has(types.KeyPrefix(types.GameKey + id))
}

// GetGameOwner returns the creator of the game
func (k Keeper) GetGameOwner(ctx sdk.Context, key string) string {
	return k.GetGame(ctx, key).Creator
}

// GetAllGame returns all game
func (k Keeper) GetAllGame(ctx sdk.Context) (msgs []types.Game) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.GameKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Game
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
