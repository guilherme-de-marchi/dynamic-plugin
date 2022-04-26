package dynamic_plugin

import (
	"fmt"
	"os"
	"plugin"
	"reflect"
	"strings"
)

type PluginManager struct {
	plugins map[string]*plugin.Plugin
}

func LoadPlugins(path string) (*PluginManager, error) {
	pls, err := loadPluginsFrom(path)
	return &PluginManager{
		plugins: pls,
	}, err
}

func (pm *PluginManager) GetPlugin(name string) (*plugin.Plugin, error) {
	p, ok := pm.plugins[name]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (pm *PluginManager) ListPlugins() []string {
	return keys(pm.plugins)
}

func (pm *PluginManager) GetSymbol(pName, symName string) (plugin.Symbol, error) {
	p, err := pm.GetPlugin(pName)
	if err != nil {
		return nil, err
	}
	s, err := p.Lookup(symName)
	if err != nil {
		return nil, ErrNotFound
	}
	return s, nil
}

func (pm *PluginManager) GetValue(pName, symName string) (reflect.Value, error) {
	s, err := pm.GetSymbol(pName, symName)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.Indirect(reflect.ValueOf(s)), nil
}

func (pm *PluginManager) GetFunc(pName, funcName string) (*Func, error) {
	v, err := pm.GetValue(pName, funcName)
	if err != nil {
		return nil, err
	}
	return NewFunc(v)
}

func (pm *PluginManager) GetStruct(pName, sName string) (*Struct, error) {
	v, err := pm.GetValue(pName, sName)
	if err != nil {
		return nil, err
	}
	return NewStruct(v)
}

func loadPluginsFrom(path string) (map[string]*plugin.Plugin, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	pls := make(map[string]*plugin.Plugin)
	for _, f := range dir {
		if f.IsDir() {
			continue
		}
		if strings.HasSuffix(f.Name(), ".so") {
			p, err := plugin.Open(fmt.Sprintf("%s/%s", path, f.Name()))
			if err != nil {
				return nil, err
			}
			pls[strings.Split(f.Name(), ".")[0]] = p
		}
	}
	return pls, nil
}
