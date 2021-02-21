package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
)

func CmdCreateGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-game [rows] [cols] [strike] [reward] [gameState] [guest] [Oplayer] [Xplayer] [nextPlayer]",
		Short: "Creates a new game",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRows := string(args[0])
			argsCols := string(args[1])
			argsStrike := string(args[2])
			argsReward := string(args[3])
			argsGameState := string(args[4])
			argsGuest := string(args[5])
			argsOplayer := string(args[6])
			argsXplayer := string(args[7])
			argsNextPlayer := string(args[8])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateGame(clientCtx.GetFromAddress().String(), string(argsRows), string(argsCols), string(argsStrike), string(argsReward), string(argsGameState), string(argsGuest), string(argsOplayer), string(argsXplayer), string(argsNextPlayer))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-game [id] [rows] [cols] [strike] [reward] [gameState] [guest] [Oplayer] [Xplayer] [nextPlayer]",
		Short: "Update a game",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsRows := string(args[1])
			argsCols := string(args[2])
			argsStrike := string(args[3])
			argsReward := string(args[4])
			argsGameState := string(args[5])
			argsGuest := string(args[6])
			argsOplayer := string(args[7])
			argsXplayer := string(args[8])
			argsNextPlayer := string(args[9])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateGame(clientCtx.GetFromAddress().String(), id, string(argsRows), string(argsCols), string(argsStrike), string(argsReward), string(argsGameState), string(argsGuest), string(argsOplayer), string(argsXplayer), string(argsNextPlayer))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-game [id] [rows] [cols] [strike] [reward] [gameState] [guest] [Oplayer] [Xplayer] [nextPlayer]",
		Short: "Delete a game by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteGame(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
