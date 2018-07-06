package compiler

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	GITCOMMIT = "HEAD"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version of concerto",
		Run: func(*cobra.Command, []string) {
			fmt.Printf("koki/concerto: %s\n", GITCOMMIT)
		},
	}
)
