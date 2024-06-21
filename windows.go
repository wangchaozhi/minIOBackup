//package main
//
//import (
//	"bufio"
//	"fmt"
//	"log"
//	"os"
//	"os/exec"
//	"strings"
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
//	allCopyCommand := fmt.Sprintf("%s cp --recursive myminio/%s ./oa-images-pro", mcPath, bucketName)
//	fmt.Println("执行:", allCopyCommand)
//	executeCommand(allCopyCommand)
//
//	watchCommand := fmt.Sprintf("%s watch --recursive myminio/%s", mcPath, bucketName)
//	fmt.Println("执行:", watchCommand)
//	cmd := exec.Command(mcPath, "watch", "--recursive", fmt.Sprintf("myminio/%s", bucketName))
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
//		cpCommand := fmt.Sprintf("%s cp myminio/%s%s ./oa-images-pro/%s%s", mcPath, bucketName, extracted, bucketName, extracted)
//		fmt.Println("执行:", cpCommand)
//		executeCommand(cpCommand)
//		fmt.Println("Copied:", images)
//	}
//}
//
//func executeCommand(command string) {
//	args := strings.Split(command, " ")
//	cmd := exec.Command(args[0], args[1:]...)
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	err := cmd.Run()
//	if err != nil {
//		log.Println(err)
//	}
//}
