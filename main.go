package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routers"
	"bluebell/settings"

	"fmt"

	"go.uber.org/zap"
)

func main() {
	// 加载配置(配置文件加载 远程加载)
	if err := settings.Init(); err != nil {
		fmt.Println("init settings failed", err)
		return
	}
	// 初始化日志  大型项目必须使用日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Println("init logger failed", err)
	}
	defer zap.L().Sync() //缓冲区日志 追加到日志文件
	zap.L().Debug("logger init success")
	// 初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("init mysql failed", err)
	}
	defer mysql.Close()
	// 初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("init redis failed", err)
	}
	defer redis.Close()

	//初始化框架编译器
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Println("init snowflake failed", err)
		return
	}
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed,err:%v\n", err)
		return
	}
	r := routers.Setup()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed,err:%v\n", err)
		return
	}
}
