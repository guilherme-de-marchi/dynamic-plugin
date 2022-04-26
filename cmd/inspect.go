package cmd

import (
	"errors"
	"fmt"
	dplugin "github.com/Guilherme-De-Marchi/dynamic-plugin"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use: "inspect {plugin_name} {func | struct | {method | field {struct_name}}} {name}",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("invalid command params")
		}
		switch args[1] {
		case "func":
		case "struct":
			break
		case "method":
		case "field":
			if len(args) < 4 {
				return errors.New("requires struct name")
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
			fmt.Println(f)
		case "struct":
			s, err := pm.GetStruct(args[0], args[2])
			if err != nil {
				return err
			}
			fmt.Println(s)
		case "method":
			s, err := pm.GetStruct(args[0], args[2])
			if err != nil {
				return err
			}
			m, err := s.GetMethod(args[3])
			if err != nil {
				return err
			}
			fmt.Println(m)
		case "field":
			s, err := pm.GetStruct(args[0], args[2])
			if err != nil {
				return err
			}
			fd, err := s.GetField(args[3])
			if err != nil {
				return err
			}
			fmt.Println(fd)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
