package main

import (
	"bufio"
	"fmt"
	dplugin "github.com/Guilherme-De-Marchi/dynamic-plugin"
	"os"
	"strings"
)

func main() {
	pm, err := dplugin.LoadPlugins("./plugins")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
terminalLoop:
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		args := strings.Split(text, " ")
		switch args[0] {
		case "list":
			if len(args) < 2 {
				fmt.Println("usage: list {plugins | methods | fields}")
				break
			}
			switch args[1] {
			case "plugins":
				fmt.Println("Plugins:")
				for i, v := range pm.ListPlugins() {
					fmt.Printf("%v : %v\n", i, v)
				}
			case "methods":
				if len(args) < 4 {
					fmt.Println("usage: list methods {plugin_name} {struct_name}")
					break
				}
				pName := args[2]
				sName := args[3]
				s, err := pm.GetStruct(pName, sName)
				if err != nil {
					fmt.Println("error: ", err)
					break
				}
				fmt.Println("Methods:")
				for i, v := range s.ListMethods() {
					fmt.Printf("%v : %v\n", i, v)
				}
			case "fields":
				if len(args) < 4 {
					fmt.Println("usage: list fields {plugin_name} {struct_name}")
					break
				}
				pName := args[2]
				sName := args[3]
				s, err := pm.GetStruct(pName, sName)
				if err != nil {
					fmt.Println("error: ", err)
					break
				}
				fmt.Println("Fields:")
				for i, v := range s.ListFields() {
					fmt.Printf("%v : %v\n", i, v)
				}
			}
		case "call":
			if len(args) < 4 {
				fmt.Println("usage: call {plugin_name} {method {struct_name} | func} {name} [args...]")
				break
			}
			pName := args[1]
			switch args[2] {
			case "method":
				if len(args) < 5 {
					fmt.Println("usage: call {plugin_name} {method {struct_name} | func} {name} [args...]")
					break
				}
				sName := args[3]
				mName := args[4]
				s, err := pm.GetStruct(pName, sName)
				if err != nil {
					fmt.Println("error: ", err)
					break
				}
				var params []string
				if len(args) >= 6 {
					params = args[5:]
				}
				fmt.Println(params) // bug
				out, err := s.Call(mName, dplugin.AnyToAny(params...)...)
				if err != nil {
					fmt.Println("error: ", err)
					break
				}
				fmt.Println("Outputs:")
				for i, v := range out {
					fmt.Printf("%v : %v\n", i, v)
				}
			}
		case "quit":
			break terminalLoop
		}
	}
	fmt.Println("done")

	//pl, err := plugin.Open("./plugins/blockchain.so")
	//if err != nil {
	//	panic(err)
	//}
	//bc, err := pl.Lookup("Bc")
	//if err != nil {
	//	panic(err)
	//}
	//s, err := dplugin.NewPStruct(reflect.Indirect(reflect.ValueOf(bc)))
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Println(s)
	//fmt.Println("fields: ", s.Fields)
	//fmt.Println("methods: ", s.Methods)
	//values, err := s.Call("ListBlocks")
	//if err != nil {
	//	panic(err)
	//}
	//for _, v := range values {
	//	fmt.Println(reflect.Indirect(v))
	//}
}
