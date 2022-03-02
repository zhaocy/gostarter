package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zhaocy/gostarter"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	app,err:=gostarter.New(viper.GetViper())
	if err != nil {
		fmt.Println("new app error")
		return
	}
	confCtx:=app.GetStartContext()
	fmt.Printf("%+v\n", confCtx)
}
