package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/umarcor/cobra"
	"github.com/umarcor/run/lib"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Exec list of tasks",
	Long:  `Exec list of tasks for the given nodes.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdOut := new(bytes.Buffer)
		cmdErr := new(bytes.Buffer)

		checkErr(lib.ExecCmd("", args[0], args[1:], nil, cmdOut, cmdErr, true))

		b, err := ioutil.ReadAll(cmdOut)
		checkErr(err)
		fmt.Println("cmdOut:\n", string(b))

		b, err = ioutil.ReadAll(cmdErr)
		checkErr(err)
		fmt.Println("cmdErr:\n", string(b))

	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	/*
		f := execCmd.PersistentFlags()
		// Bind the full flag set to the configuration
		err := v.BindPFlags(f)
		checkErr(err)
	*/
}
