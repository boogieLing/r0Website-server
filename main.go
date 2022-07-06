/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 提供网站的后端网络服务，服务器启动入口
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/3 18:16
 */
package main

import (
	"context"
	"fmt"
	"os"
	"r0Website-server/core"
	"r0Website-server/global"
	"r0Website-server/initialize"
	"strconv"
	"strings"
)

// using env:   export GIN_MODE=release
func main() {
	// 解析参数
	if len(os.Args) > 1 {
		for idx, arg := range os.Args {
			fmt.Println("参数"+strconv.Itoa(idx)+" : ", arg)
		}
		arg := strings.Split(os.Args[1], "=")
		if len(arg) >= 2 && arg[0] == "--config" {
			// 指定yaml文件的路径
			global.Config = initialize.InitProdConfig(arg[1])
		}
		/*if len(arg) >= 4 && arg[2] == "--lv" {
			// 设置项目级别
			global.ProjLevel = initialize.InitProjLevel(arg[2])
		}*/
	} else {
		global.Config = initialize.InitDevConfig()
		global.ProjLevel = initialize.InitProjLevel("DEBUG")
	}
	initialize.InitLogger()
	var ctx context.Context
	ctx, global.ClientEngine = initialize.InitDB()
	global.DBEngine = global.ClientEngine.Database(global.Config.Mongo.DB)
	// 必须关闭 但是defer
	defer func() {
		if err := global.ClientEngine.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	initialize.InitUtils()
	core.RunWindowsServer()
}
