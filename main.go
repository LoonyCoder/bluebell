package main

import (
	"BlueBell/controller"
	"BlueBell/dao/mysql"
	"BlueBell/dao/redis"
	"BlueBell/logger"
	"BlueBell/pkg/snowflake"
	"BlueBell/routers"
	"BlueBell/settings"
	"fmt"
)

func main() {
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed,err:%v\n", err)
		return
	}
	// 加载配置
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	// 加载配置
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	// 加载配置
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed , err:%v\n", err)
		return
	}

	//初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed , err:%v\n", err)
		return
	}

	// 注册路由
	engine := routers.SetupRouter(settings.Conf.Mode)
	err := engine.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed , err:%v\n", err)
		return
	}

	//go func() {
	//	// 开启一个goroutine启动服务
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen:%s\n", err)
	//	}
	//}()
}
