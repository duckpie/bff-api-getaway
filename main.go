package main

import (
	"github.com/spf13/cobra"
	"github.com/wrs-news/bff-api-getaway/cmd"
)

func main() {
	cobra.CheckErr(cmd.NewRootCmd().Execute())
}
