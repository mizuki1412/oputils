package mod_api_test

import (
	"github.com/mizuki1412/go-core-kit/v2/class/exception"
	"github.com/mizuki1412/go-core-kit/v2/library/filekit"
	"github.com/mizuki1412/go-core-kit/v2/library/jsonkit"
	"github.com/mizuki1412/go-core-kit/v2/service/configkit"
)

type Unit struct {
	Url    string         `json:"url"`
	Method string         `json:"method"`
	Params map[string]any `json:"params"`
	Body   map[string]any `json:"body"`
}

func Run() {
	configPath := configkit.GetString("file")
	json, err := filekit.ReadString(configPath)
	if err != nil {
		panic(exception.New(err.Error()))
	}
	var units []Unit
	err = jsonkit.ParseObj(json, units)
	if err != nil {
		panic(exception.New(err.Error()))
	}
	for _, e := range units {
		println(e)
	}
}
