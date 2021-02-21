package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/misaakidis/cosmos-triliza/x/cosmostriliza/types"
	"github.com/spf13/cobra"
)

func CmdListGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-game",
		Short: "list all game",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGameRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GameAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-game [id]",
		Short: "shows a game",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetGameRequest{
				Id: args[0],
			}

			res, err := queryClient.Game(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
