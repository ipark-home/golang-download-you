package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/kkdai/youtube"
)

func downFile(urlYoutube string, fileName string) {
	log.Println("download to file=", fileName)
	y := NewYoutube(true)
	y.DecodeURL(urlYoutube)
	y.StartDownload(fileName)
}
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	fmt.Println("...start...")

	fileName := "f:/files.new.txt"
	lines, err := readLines(fileName)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {
		fmt.Println(i, line)
		args := strings.Fields(line)
		downFile(args[0], args[1])
	}

	fmt.Println("...finish...")

}

// func main() {
// 	urlYoutube, _ := filepath.Abs(os.Args[1])
// 	currentFile, _ := filepath.Abs(os.Args[2])
// 	log.Println("download to file=", currentFile)
// 	y := NewYoutube(true)
// 	y.DecodeURL(urlYoutube)
// 	y.StartDownload(currentFile)
// }
