package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// 定义命令行参数
	endpoints := flag.String("endpoints", "localhost:2379", "etcd endpoints, comma-separated")
	username := flag.String("username", "", "etcd username (optional)")
	password := flag.String("password", "", "etcd password (optional)")
	flag.Parse()

	// 将逗号分隔的端点转换为切片
	endpointList := strings.Split(*endpoints, ",")

	// 配置etcd客户端
	cfg := clientv3.Config{
		Endpoints:   endpointList,
		DialTimeout: 5 * time.Second,
	}

	// 仅在提供了用户名和密码时设置认证
	if *username != "" && *password != "" {
		cfg.Username = *username
		cfg.Password = *password
	}

	// 创建etcd客户端
	cli, err := clientv3.New(cfg)
	if err != nil {
		log.Fatalf("创建etcd客户端失败: %v", err)
	}
	defer cli.Close()

	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 跟踪检查结果
	var checkFailed bool

	// 检查集群健康状态
	fmt.Println("=== 检查etcd集群状态 ===")
	if err := checkClusterHealth(ctx, cli, endpointList); err != nil {
		log.Printf("集群健康检查失败: %v", err)
		checkFailed = true
	}

	// 获取集群成员列表
	if err := listClusterMembers(ctx, cli); err != nil {
		log.Printf("获取成员列表失败: %v", err)
		checkFailed = true
	}

	// 获取领导者信息
	if err := getLeaderInfo(ctx, cli); err != nil {
		log.Printf("获取领导者信息失败: %v", err)
		checkFailed = true
	}

	// 根据检查结果打印最终状态
	if checkFailed {
		fmt.Println("failed")
		os.Exit(2)
	} else {
		fmt.Println("success")
	}
}

// 检查集群健康状态
func checkClusterHealth(ctx context.Context, cli *clientv3.Client, endpoints []string) error {
	resp, err := cli.Cluster.MemberList(ctx)
	if err != nil {
		return fmt.Errorf("获取成员列表失败: %v", err)
	}

	for _, member := range resp.Members {
		// 为每个成员的客户端URL创建临时客户端以检查健康
		for _, url := range member.ClientURLs {
			cfg := clientv3.Config{
				Endpoints:   []string{url},
				DialTimeout: 5 * time.Second,
			}
			// 继承主客户端的认证配置
			if cli.Username != "" && cli.Password != "" {
				cfg.Username = cli.Username
				cfg.Password = cli.Password
			}
			tempCli, err := clientv3.New(cfg)
			if err != nil {
				fmt.Printf("成员 %s (%d) 创建客户端失败: %v\n", member.Name, member.ID, err)
				return err
			}
			defer tempCli.Close()

			// 执行简单的Get操作以检查健康
			_, err = tempCli.Get(ctx, "health-check", clientv3.WithLimit(1))
			if err != nil {
				fmt.Printf("成员 %s (%d) 健康检查失败 (URL: %s): %v\n", member.Name, member.ID, url, err)
				return err
			}
			fmt.Printf("成员 %s (%d) 健康状态: 正常 (URL: %s)\n", member.Name, member.ID, url)
		}
	}
	return nil
}

// 列出集群成员
func listClusterMembers(ctx context.Context, cli *clientv3.Client) error {
	resp, err := cli.Cluster.MemberList(ctx)
	if err != nil {
		return fmt.Errorf("获取成员列表失败: %v", err)
	}

	fmt.Println("\n=== 集群成员列表 ===")
	for _, member := range resp.Members {
		fmt.Printf("成员ID: %d\n", member.ID)
		fmt.Printf("名称: %s\n", member.Name)
		fmt.Printf("端点: %v\n", member.ClientURLs)
		fmt.Printf("是否为学习者: %v\n", member.IsLearner)
		fmt.Println("---")
	}
	return nil
}

// 获取领导者信息
func getLeaderInfo(ctx context.Context, cli *clientv3.Client) error {
	resp, err := cli.Status(ctx, cli.Endpoints()[0])
	if err != nil {
		return fmt.Errorf("获取状态失败: %v", err)
	}

	fmt.Println("\n=== 领导者信息 ===")
	fmt.Printf("领导者ID: %d\n", resp.Leader)
	fmt.Printf("版本: %s\n", resp.Version)
	fmt.Printf("数据库大小: %d bytes\n", resp.DbSize)
	fmt.Printf("Raft任期: %d\n", resp.RaftTerm)
	return nil
}
