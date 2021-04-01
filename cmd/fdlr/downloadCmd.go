package main

import (
	"runtime"

	"github.com/Imputes/fdlr/internal/errorHandle"
	"github.com/Imputes/fdlr/internal/tool"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var conc int

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().IntVarP(&conc, "goroutines count", "c", runtime.NumCPU(), "default is the number of CPU cores for the current environment")
}

var downloadCmd = &cobra.Command{
	Use:     "download",
	Short:   "downloads a file from URL or file name",
	Example: `fdlr download [-c=goroutines_count] URL`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		errorHandle.ExitWithError(download(args[0]))
	},
}

func download(URL string) error {
	_, err := tool.GetFolderFrom(URL)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
