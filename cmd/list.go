package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"

	"github.com/Guilherme-De-Marchi/dynamic-plugin"
)

var plugin string
var structure string

var listCmd = &cobra.Command{
	Use: "list {plugins | {methods | fields} plugin_name struct_name}",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a list target")
		}
		switch args[0] {
		case "plugins":
			break
		case "methods":
		case "fields":

			if len(args) < 3 {
				return errors.New("requires the plugin_name and struct_name")
			}
		default:
			return errors.New("invalid list target")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		pm, err := dypl.LoadPlugins("./plugins")
		if err != nil {
			return err
		}
		switch args[0] {
		case "plugins":
			fmt.Println(pm.ListPlugins())
		case "methods":
			st, err := pm.GetStruct(args[1], args[2])
			if err != nil {
				return err
			}
			fmt.Println(st.ListMethods())
		case "fields":
			st, err := pm.GetStruct(args[1], args[2])
			if err != nil {
				return err
			}
			fmt.Println(st.ListFields())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
