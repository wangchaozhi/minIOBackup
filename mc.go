//package main
//
//import (
//	"bufio"
//	"bytes"
//	"fmt"
//	"log"
//	"os/exec"
//	"strings"
//)
//
//var bucketName string
//
//func main() {
//	// 从环境变量获取 bucketName
//	bucketName = "oa-images"
//	allCopyCommand := fmt.Sprintf("mc cp --recursive  myminio/%s ./oa-images-pro", bucketName)
//	fmt.Println("执行:", allCopyCommand)
//	go executeCommand(allCopyCommand)
//
//	watchCommand := fmt.Sprintf("mc watch --recursive myminio/%s", bucketName)
//	fmt.Println("执行:", watchCommand)
//	out, err := exec.Command("sh", "-c", watchCommand).Output()
//	if err != nil {
//		log.Fatal(err)
//	}
//	reader := bufio.NewReader(bytes.NewReader(out))
//	for {
//		line, _, err := reader.ReadLine()
//		if err != nil {
//			break
//		}
//		if strings.Contains(string(line), "ObjectCreated:Put") {
//			fmt.Println(string(line))
//			startIndex := strings.Index(string(line), bucketName)
//			var extracted string
//			if startIndex != -1 {
//				extracted = string(line)[startIndex+len(bucketName):]
//			}
//			if extracted == "" {
//				fmt.Println("路径为null")
//				break
//			}
//			images := strings.Replace(extracted, "/", "", -1)
//			fmt.Println(images)
//			cpCommand := fmt.Sprintf("mc cp myminio/%s%s ./oa-images-pro/%s%s", bucketName, extracted, bucketName, extracted)
//			fmt.Println("执行:", cpCommand)
//			executeCommandAndWait(cpCommand)
//			fmt.Println("Copied:", images)
//		}
//	}
//}
//
//func executeCommand(command string) {
//	cmd := exec.Command("sh", "-c", command)
//	_, err := cmd.Output()
//	if err != nil {
//		log.Println(err)
//	}
//}
//
//func executeCommandAndWait(command string) {
//	cmd := exec.Command("sh", "-c", command)
//	err := cmd.Run()
//	if err != nil {
//		log.Println(err)
//	}
//}
