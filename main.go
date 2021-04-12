package main

import (
	"fmt"
	"github.com/mengjunwei/go_api_project/services/linux_system_source"
	"os"
	"runtime"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	proc_util "github.com/mengjunwei/go-utils/proc-util"

	"github.com/mengjunwei/go_api_project/logger"
	_ "github.com/mengjunwei/go_api_project/routers"
)

func main() {
	goMaxProcs := beego.AppConfig.DefaultInt("GoMaxProcs", 8)
	runtime.GOMAXPROCS(goMaxProcs)

	// 跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
	}))

	if err := logger.SetLog(); err != nil {
		fmt.Sprintln(err.Error())
		os.Exit(1)
	}

	go func() {
		beego.Run()
	}()
	go func() {
		linux_system_source.ReadMemList()
	}()
	proc_util.WaitForSigterm()

}
