package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
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
	GameState  string       `json:"gameState"`
	Guest      string       `json:"guest"`
	Oplayer    string       `json:"Oplayer"`
	Xplayer    string       `json:"Xplayer"`
	NextPlayer string       `json:"nextPlayer"`
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

		parsedRows := req.Rows

		parsedCols := req.Cols

		parsedStrike := req.Strike

		parsedReward := req.Reward

		parsedGameState := req.GameState

		parsedGuest := req.Guest

		parsedOplayer := req.Oplayer

		parsedXplayer := req.Xplayer

		parsedNextPlayer := req.NextPlayer

		msg := types.NewMsgCreateGame(
			req.Creator,
			parsedRows,
			parsedCols,
			parsedStrike,
			parsedReward,
			parsedGameState,
			parsedGuest,
			parsedOplayer,
			parsedXplayer,
			parsedNextPlayer,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateGameRequest struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	Creator    string       `json:"creator"`
	Rows       string       `json:"rows"`
	Cols       string       `json:"cols"`
	Strike     string       `json:"strike"`
	Reward     string       `json:"reward"`
	GameState  string       `json:"gameState"`
	Guest      string       `json:"guest"`
	Oplayer    string       `json:"Oplayer"`
	Xplayer    string       `json:"Xplayer"`
	NextPlayer string       `json:"nextPlayer"`
}

func updateGameHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var req updateGameRequest
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

		parsedRows := req.Rows

		parsedCols := req.Cols

		parsedStrike := req.Strike

		parsedReward := req.Reward

		parsedGameState := req.GameState

		parsedGuest := req.Guest

		parsedOplayer := req.Oplayer

		parsedXplayer := req.Xplayer

		parsedNextPlayer := req.NextPlayer

		msg := types.NewMsgUpdateGame(
			req.Creator,
			id,
			parsedRows,
			parsedCols,
			parsedStrike,
			parsedReward,
			parsedGameState,
			parsedGuest,
			parsedOplayer,
			parsedXplayer,
			parsedNextPlayer,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteGameRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deleteGameHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var req deleteGameRequest
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

		msg := types.NewMsgDeleteGame(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
