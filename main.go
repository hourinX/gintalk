package main

import (
	"fmt"
	"gin-online-chat-backend/router"
	"gin-online-chat-backend/systems"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := systems.LoadConfig("config/config.yml"); err != nil {
		log.Fatalf("配置文件加载失败: %v", err)
	}

	if err := systems.InitDatabase(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	defer systems.CloseDatabase()

	if err := systems.InitRedis(); err != nil {
		log.Fatalf("Redis初始化失败: %v", err)
	}
	defer systems.CloseRedis()

	if err := systems.InitRabbitMQ(); err != nil {
		log.Fatalf("RabbitMQ 初始化失败: %v", err)
	}
	defer systems.CloseRabbitMQ()

	if err := systems.InitElasticsearch(); err != nil {
		log.Fatalf("ElasticSearch初始化失败: %v", err)
	} else {
		// 创建用户索引
		//searchService := search.NewSearchService()
		//if err = searchService.CreateUserIndex(); err != nil {
		//	log.Fatalf("Warning: Failed to create user index: %v", err)
		//}
	}

	r := router.SetupRouter()
	config := systems.GetConfig()
	addr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	log.Printf("服务启动中,正在监听地址: %s", addr)

	go func() {
		if err := r.Run(addr); err != nil {
			log.Fatalf("协程异步启动失败: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("服务正在关闭")
}
