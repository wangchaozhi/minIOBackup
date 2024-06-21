//package main
//
//import (
//	"bufio"
//	"fmt"
//	"log"
//	"os"
//	"os/exec"
//	"strings"
//	"sync"
//)
//
//var bucketName = "oa-images"
//
//func main() {
//	currentDir, err := os.Getwd()
//	if err != nil {
//		log.Fatalf("Failed to get current directory: %v", err)
//	}
//	mcPath := currentDir + "\\mc.exe"
//
//	allCopyCommand := []string{mcPath, "cp", "--recursive", fmt.Sprintf("myminio/%s", bucketName), "./oa-images-pro"}
//	fmt.Println("执行:", strings.Join(allCopyCommand, " "))
//
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		executeCommand(allCopyCommand)
//	}()
//
//	watchCommand := []string{mcPath, "watch", "--recursive", fmt.Sprintf("myminio/%s", bucketName)}
//	fmt.Println("执行:", strings.Join(watchCommand, " "))
//	cmd := exec.Command(watchCommand[0], watchCommand[1:]...)
//	cmd.Stderr = os.Stderr
//	stdout, _ := cmd.StdoutPipe()
//	if err := cmd.Start(); err != nil {
//		log.Fatal(err)
//	}
//
//	reader := bufio.NewReader(stdout)
//	for {
//		line, _, err := reader.ReadLine()
//		if err != nil {
//			break
//		}
//		processLine(string(line), mcPath)
//	}
//
//	if err := cmd.Wait(); err != nil {
//		log.Println(err)
//	}
//
//	wg.Wait() // 等待所有goroutines完成
//}
//
//func processLine(line, mcPath string) {
//	if strings.Contains(line, "ObjectCreated:Put") {
//		fmt.Println(line)
//		startIndex := strings.Index(line, bucketName)
//		var extracted string
//		if startIndex != -1 {
//			extracted = line[startIndex+len(bucketName):]
//		}
//		if extracted == "" {
//			fmt.Println("路径为null")
//			return
//		}
//		images := strings.Replace(extracted, "/", "", -1)
//		fmt.Println(images)
//		cpCommand := []string{mcPath, "cp", fmt.Sprintf("myminio/%s%s", bucketName, extracted), fmt.Sprintf("./oa-images-pro/%s%s", bucketName, extracted)}
//		fmt.Println("执行:", strings.Join(cpCommand, " "))
//		executeCommand(cpCommand)
//		fmt.Println("Copied:", images)
//	}
//}
//
//func executeCommand(command []string) {
//	cmd := exec.Command(command[0], command[1:]...)
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	err := cmd.Run()
//	if err != nil {
//		log.Println(err)
//	}
//}
