package cmd

import (
	"errors"
	"fmt"
	dplugin "github.com/Guilherme-De-Marchi/dynamic-plugin"
	"github.com/spf13/cobra"
)

var callCmd = &cobra.Command{
	Use:   "call {plugin_name} {func | method {struct_name}} {name} [args...]",
	Short: "call a method",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("invalid command params")
		}
		switch args[1] {
		case "func":
			break
		case "method":
			if len(args) < 4 {
				return errors.New("required struct name")
			}
		default:
			return errors.New("invalid target")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		pm, err := dplugin.LoadPlugins("./plugins")
		if err != nil {
			return err
		}
		switch args[1] {
		case "func":
			f, err := pm.GetFunc(args[0], args[2])
			if err != nil {
				return err
			}
			var in []string
			if len(args) > 3 {
				in = args[3:]
			}
			fmt.Println(f.Call(dplugin.AnyToValue(in)...))
		case "method":
			s, err := pm.GetStruct(args[0], args[1])
			if err != nil {
				return err
			}
			var in []string
			if len(args) > 3 {
				in = args[3:]
			}
			out, err := s.Call(args[2], dplugin.AnyToAny(in...)...)
			if err != nil {
				return err
			}
			fmt.Println(out)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(callCmd)
}
