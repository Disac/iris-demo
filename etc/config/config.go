package config

import (
	"flag"
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

var (
	Env  = os.Getenv("ENVIRON")
	Port = flag.String("port", "8000", "http listen port")
	Path = flag.String("config", "", "config file path.")
)
var conf *config

func init() {
	flag.Parse()
	conf = &config{}
	err := conf.Init()
	if err != nil {
		panic(err)
	}
}

type config struct {
	tree *toml.Tree
}

func (conf *config) Init() (err error) {
	if Env == "" {
		Env = "develop"
	}
	path := *Path
	if path == "" {
		path = os.Getenv("GOCONF")
	}

	path += "/demo." + Env + ".toml"
	conf.tree, err = toml.LoadFile(path)
	if err != nil {
		return
	}
	return
}

func Val(key string) string {
	fmt.Println(conf.tree.Get(key))
	return conf.tree.Get(key).(string)
}

func Int(key string) int {
	return int(conf.tree.Get(key).(int64))
}

func Int64(key string) int64 {
	return conf.tree.Get(key).(int64)
}

func Map(key string) map[string]interface{} {
	tree := conf.tree.Get(key).(*toml.Tree)
	return tree.ToMap()
}
