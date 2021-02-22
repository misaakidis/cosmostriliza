package cli

import (
	"strconv"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
)

func CmdCreateGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-game [rows] [cols] [strike] [reward]",
		Short: "Creates a new game",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRows, _ := strconv.ParseUint(args[0], 10, 32)
			argsCols, _ := strconv.ParseUint(args[1], 10, 32)
			argsStrike, _ := strconv.ParseUint(args[2], 10, 32)
			argsReward, _ := strconv.ParseUint(args[3], 10, 32)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateGame(clientCtx.GetFromAddress().String(), uint32(argsRows), uint32(argsCols), uint32(argsStrike), uint32(argsReward))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}


func CmdJoinGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "join-game [id]",
		Short: "Joins an existing game with the given id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsId := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgJoinGame(clientCtx.GetFromAddress().String(), argsId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
