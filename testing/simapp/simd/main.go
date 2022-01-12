package main

import (
	"os"

	"github.com/reapchain/cosmos-sdk/server"
	svrcmd "github.com/reapchain/cosmos-sdk/server/cmd"
	"github.com/cosmos/ibc-go/v2/testing/simapp"
	"github.com/cosmos/ibc-go/v2/testing/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
