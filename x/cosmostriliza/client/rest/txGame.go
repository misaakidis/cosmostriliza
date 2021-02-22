package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createGameRequest struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	Creator    string       `json:"creator"`
	Rows       string       `json:"rows"`
	Cols       string       `json:"cols"`
	Strike     string       `json:"strike"`
	Reward     string       `json:"reward"`
}

func createGameHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createGameRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedRows, _ := strconv.ParseUint(req.Rows, 10, 32)
		parsedCols, _ := strconv.ParseUint(req.Cols, 10, 32)
		parsedStrike, _ := strconv.ParseUint(req.Strike, 10, 32)
		parsedReward, _ := strconv.ParseUint(req.Reward, 10, 32)

		msg := types.NewMsgCreateGame(
			req.Creator,
			uint32(parsedRows),
			uint32(parsedCols),
			uint32(parsedStrike),
			uint32(parsedReward),
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
