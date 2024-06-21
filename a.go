//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os/exec"
//	"strings"
//)
//
//var bucketName string = "YOUR_BUCKET_NAME" // replace with your actual bucket name or get it from a configuration file
//
//func main() {
//	allCopyCommand := fmt.Sprintf("mc cp --recursive  myminio/%s ./oa-images-pro", bucketName)
//	fmt.Println("Executing", allCopyCommand)
//	go executeCommand(allCopyCommand)
//
//	watchCommand := fmt.Sprintf("mc watch --recursive myminio/%s", bucketName)
//	fmt.Println("Executing", watchCommand)
//	executeWatchCommand(watchCommand)
//}
//
//func executeCommand(command string) {
//	cmd := exec.Command("sh", "-c", command) // Assuming you're on a Unix-like system
//	out, _ := cmd.CombinedOutput()           // Error handling omitted for brevity
//	fmt.Println(string(out))
//}
//
//func executeWatchCommand(command string) {
//	cmd := exec.Command("sh", "-c", command)
//	stdout, _ := cmd.StdoutPipe()
//	cmd.Start()
//
//	scanner := bufio.NewScanner(stdout)
//	for scanner.Scan() {
//		line := scanner.Text()
//		if strings.Contains(line, "ObjectCreated:Put") {
//			fmt.Println(line)
//
//			startIndex := strings.Index(line, bucketName)
//			var extracted string
//			if startIndex != -1 {
//				extracted = line[startIndex+len(bucketName):]
//			}
//
//			if extracted == "" {
//				fmt.Println("Path is null")
//				break
//			}
//
//			images := strings.Replace(extracted, "/", "", -1)
//			fmt.Println(images)
//
//			cpCommand := fmt.Sprintf("mc cp myminio/%s%s ./oa-images-pro/%s%s", bucketName, extracted, bucketName, extracted)
//			fmt.Println("Executing", cpCommand)
//			executeCommand(cpCommand)
//		}
//	}
//
//	cmd.Wait()
//}
