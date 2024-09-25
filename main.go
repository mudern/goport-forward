package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type Config struct {
	Source []string `json:"source"`
	Target []string `json:"target"`
}

// 默认配置
var defaultConfig = Config{
	Source: []string{"127.0.0.1:80"},
	Target: []string{"127.0.0.1:80"},
}

func handleConnection(src net.Conn, target string) {
	defer src.Close()

	// 连接到目标地址
	dest, err := net.Dial("tcp", target)
	if err != nil {
		log.Printf("Failed to connect to target %s: %v", target, err)
		return
	}
	defer dest.Close()

	// 同步数据传输
	go io.Copy(dest, src) // 从源连接到目标连接
	io.Copy(src, dest)    // 从目标连接到源连接
}

func startForwarding(source string, target string) {
	listener, err := net.Listen("tcp", source)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", source, err)
	}
	defer listener.Close()

	log.Printf("Listening on %s and forwarding to %s", source, target)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn, target) // 为每个连接启动一个 goroutine
	}
}

func loadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		// 如果文件不存在，则创建默认的 config.json
		log.Printf("Config file not found, creating default config file: %s", filename)
		err = createDefaultConfig(filename)
		if err != nil {
			return config, err
		}
		return defaultConfig, nil
	} else if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	return config, err
}

func createDefaultConfig(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将默认配置编码为JSON
	bytes, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return err
	}

	// 写入文件
	_, err = file.Write(bytes)
	return err
}

func main() {
	// 加载配置文件
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	// 确保 source 和 target 数量一致
	if len(config.Source) != len(config.Target) {
		log.Fatalf("Source and target arrays must have the same length")
	}

	// 启动每对 source-target 的转发
	for i := range config.Source {
		go startForwarding(config.Source[i], config.Target[i])
	}

	// 保持主线程运行
	select {}
}
