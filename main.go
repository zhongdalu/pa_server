//@Auth:zdl
package main

import (
	"com.deer.pa_server/server"
	"gitee.com/sky_big/workmod"
	"github.com/gogf/gf/g/os/glog"
)

func main(){
	glog.Println("server start")
	workmod.Init()
	server.Run()
	select{}
}
