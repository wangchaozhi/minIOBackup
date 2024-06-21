package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time" // 引入 time 包来处理延时
)

var bucketName = "oa-images-pro"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in main:", r)
		}
	}()

	for {
		if err := watchMC(); err != nil {
			log.Printf("Error watching mc: %v", err)
			time.Sleep(6 * time.Second) // 在重试之前等待 6 秒
			log.Printf("reconnect......")
		} else {
			log.Printf("mc正常退出")
			break // 如果 watchMC 正常退出，则停止循环
		}
	}
}

func watchMC() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Failed to get current directory: %v", err)
	}
	mcPath := currentDir + "/mc" // 修改了路径分隔符和文件名

	watchCommand := []string{mcPath, "watch", "--events=put", "--recursive", fmt.Sprintf("myminio/%s", bucketName)}
	cmd := exec.Command(watchCommand[0], watchCommand[1:]...)
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("StdoutPipe error: %v", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Start error: %v", err)
	}
	// 在这里记录连接成功的消息
	log.Printf("mc watch command started successfully.")

	reader := bufio.NewReader(stdout)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				log.Println("EOF received, stopping watch")
				break // 正常结束循环
			}
			return fmt.Errorf("ReadLine error: %v", err)
		}
		processLine(string(line), mcPath)
	}
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("Wait error: %v", err)
	}

	return nil
}

func processLine(line, mcPath string) {
	if strings.Contains(line, "ObjectCreated:Put") {
		fmt.Println(line)
		startIndex := strings.Index(line, bucketName)
		var extracted string
		if startIndex != -1 {
			extracted = line[startIndex+len(bucketName):]
		}
		if extracted == "" {
			fmt.Println("路径为null")
			return
		}
		if strings.Contains(extracted, "条形码") { // 检查路径是否包含“条形码”
			return // 如果包含，直接返回，不做其他处理
		}
		images := strings.Replace(extracted, "/", "", -1)
		fmt.Println(images)
		cpCommand := []string{mcPath, "cp", fmt.Sprintf("myminio/%s%s", bucketName, extracted), fmt.Sprintf("./oa-images%s", extracted)}
		fmt.Println("执行:", strings.Join(cpCommand, " "))
		executeCommand(cpCommand)
		fmt.Println("Copied:", images)
	}
}

func executeCommand(command []string) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
