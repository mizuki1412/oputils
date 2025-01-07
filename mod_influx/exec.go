package mod_influx

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mizuki1412/go-core-kit/v2/class/exception"
	"github.com/mizuki1412/go-core-kit/v2/library/filekit"
	"github.com/mizuki1412/go-core-kit/v2/library/timekit"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"strings"
)

func Exec() {
	client := influxdb2.NewClient(viper.GetString("url"), viper.GetString("token"))
	if client == nil {
		panic(exception.New("client nil"))
	}
	//writeApi := client.WriteAPIBlocking(viper.GetString("org"), viper.GetString("bucket"))
	switch viper.GetInt32("type") {
	case 0, 1:
		queryApi := client.QueryAPI(viper.GetString("org"))
		flux, _ := filekit.ReadString(viper.GetString("flux"))
		if flux == "" {
			panic(exception.New("flux nil"))
		}
		res, err := queryApi.Query(context.Background(), flux)
		if err != nil {
			panic(exception.New(err.Error()))
		}
		values := ""
		var keys []string
		for res.Next() {
			//if res.TableChanged() {
			//	fmt.Printf("table: %s\n", res.TableMetadata().String())
			//}
			if values == "" && len(res.Record().Values()) > 0 {
				for k := range res.Record().Values() {
					keys = append(keys, k)
				}
				values += strings.Join(keys, ", ") + "\n"
			}
			var vs []string
			// Access data
			for _, k := range keys {
				vs = append(vs, cast.ToString(res.Record().Values()[k]))
			}
			values += strings.Join(vs, ", ") + "\n"
		}
		if viper.GetInt32("type") == 0 {
			fmt.Println(values)
		} else {
			_ = filekit.WriteFile(viper.GetString("csv"), []byte(values))
		}
	case 10:
		// todo
		deleteAPI := client.DeleteAPI()
		err := deleteAPI.DeleteWithName(context.Background(), viper.GetString("org"), viper.GetString("bucket"), timekit.ParseD(viper.GetString("dstart")), timekit.ParseD(viper.GetString("dend")), "")
		if err != nil {
			panic(exception.New(err.Error()))
		}
	}
	client.Close()
}
