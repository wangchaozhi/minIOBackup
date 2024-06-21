package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var bucketName = "oa-images-pro"

func main() {
	defer recoverFromPanic("main")
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	mcPath := currentDir + "/mc" // 修改了路径分隔符和文件名

	watchCommand := []string{mcPath, "watch", "--event=put", "--recursive", fmt.Sprintf("myminio/%s", bucketName)}
	fmt.Println("执行:", strings.Join(watchCommand, " "))
	cmd := exec.Command(watchCommand[0], watchCommand[1:]...)
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(stdout)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}
		processLine(string(line), mcPath)
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err)
	}
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
	defer recoverFromPanic("executeCommand")
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func recoverFromPanic(functionName string) {
	if r := recover(); r != nil {
		log.Printf("Recovered in %s: %v\n", functionName, r)
		// 这里可以添加更多的错误处理逻辑
	}
}