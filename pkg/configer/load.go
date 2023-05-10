package configer

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var LoadType = map[string]func(path, name string, config any){
	".yaml":       LoadLocalYaml,
	".yml":        LoadLocalYaml,
	".toml":       LoadLocalToml,
	".json":       LoadLocalJson,
	".properties": LoadLocalProps,
	".props":      LoadLocalProps,
	".prop":       LoadLocalProps,
}

func MustLoadLocal(path string, config any) {
	LoadType[filepath.Ext(path)](filepath.Dir(path), filepath.Base(path), config)
}

func LoadLocalYaml(path, name string, config any) {
	LoadLocalWithTypeName(path, name, "yaml", config)
}

func LoadLocalToml(path, name string, config any) {
	LoadLocalWithTypeName(path, name, "toml", config)
}

func LoadLocalProps(path, name string, config any) {
	LoadLocalWithTypeName(path, name, "prop", config)
}

func LoadLocalJson(path, name string, config any) {
	LoadLocalWithTypeName(path, name, "json", config)
}

func LoadLocalWithTypeName(path, name, typeName string, config any) {
	execPath, _ := os.Executable()
	configFile := filepath.Join(filepath.Dir(execPath), path)
	v := viper.New()
	v.AddConfigPath(configFile)
	v.SetConfigName(name)
	v.SetConfigType(typeName)
	err := v.ReadInConfig()
	if err != nil {
		v.AddConfigPath(path)
		err = v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config files: %s \n", err))
		}
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config files changed:", e.Name)
		if err = v.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&config); err != nil {
		panic(err)
	}
}
