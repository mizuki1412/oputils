package mod_api

import (
	"fmt"
	"github.com/mizuki1412/go-core-kit/v2/class/exception"
	"github.com/mizuki1412/go-core-kit/v2/library/filekit"
	"github.com/mizuki1412/go-core-kit/v2/library/httpkit"
	"github.com/mizuki1412/go-core-kit/v2/library/jsonkit"
	"github.com/mizuki1412/go-core-kit/v2/service/configkit"
	"github.com/mizuki1412/go-core-kit/v2/service/logkit"

	"github.com/spf13/cast"
	"time"
)

type Unit struct {
	Title  string            `json:"title"`
	Url    string            `json:"url"`
	Method string            `json:"method"`
	Query  map[string]string `json:"query"`
	Form   map[string]string `json:"form"`
	Json   map[string]any    `json:"json"`
}

func Run() {
	configPath := configkit.GetString("file")
	json, err := filekit.ReadString(configPath)
	if err != nil {
		panic(exception.New(err.Error()))
	}
	dest := configkit.GetString("dest")
	err = filekit.CheckFilePath(dest)
	if err != nil {
		panic(exception.New(err.Error()))
	}
	filekit.WriteFile(dest, []byte(""))

	var units []Unit
	err = jsonkit.ParseObj(json, &units)
	if err != nil {
		panic(exception.New(err.Error()))
	}
	for _, e := range units {
		req := httpkit.Req{
			Method: e.Method,
			Url:    e.Url,
		}
		if e.Query != nil {
			req.QueryData = e.Query
		}
		if e.Form != nil {
			req.FormData = e.Form
		}
		if e.Json != nil {
			req.JsonData = e.Json
		}
		start := time.Now()
		res, code := httpkit.Request(req)
		cost := time.Since(start)
		length := len(res)
		if length > 100 {
			res = res[:100]
		}
		result := fmt.Sprintf(`
title: %s
url: %s
cost: %f s
code: %d
length: %d
result-sample: %s ...
`, e.Title, e.Url, cast.ToFloat64(cost.Milliseconds())/1000.0, code, length, res)
		logkit.Info(result)
		err = filekit.WriteFileAppend(dest, []byte(result))
		if err != nil {
			panic(exception.New(err.Error()))
		}
	}
}
