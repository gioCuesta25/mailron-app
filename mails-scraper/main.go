package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

func main() {
	f, perr := os.Create("cpu.pprof")
	if perr != nil {
		log.Fatal(perr)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	filesPath := "C:/Users/Giovanni/Downloads/enron_mail_20110402/maildir"
	files, err := os.ReadDir(filesPath)
	checkError(err)

	mails := []map[string]string{}

	for _, f := range files {
		if f.IsDir() {
			path := filesPath + fmt.Sprintf("/%s", f.Name())
			files, err = os.ReadDir(path)
			checkError(err)
			for _, j := range files {
				if j.IsDir() && j.Name() == "_sent_mail" {
					path := path + fmt.Sprintf("/%s", j.Name())
					files, err = os.ReadDir(path)
					checkError(err)
					for _, k := range files {
						if !k.IsDir() {
							path := path + fmt.Sprintf("/%s", k.Name())
							fmt.Println(path)
							info := getMailInfo(path)
							mails = append(mails, info)
						}
					}
				}
			}
		}
	}

	records := map[string][]map[string]string{}

	records["records"] = mails

	jsonFile, err := json.MarshalIndent(records, "", " ")

	checkError(err)

	err = os.WriteFile("mails.json", jsonFile, 0644)

	checkError(err)

	fmt.Println(len(mails))
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readSpecificLine(fileName string, n int) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("invalid request: line %d", n)
	}
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	bf := bufio.NewReader(f)
	var line string
	for lnum := 0; lnum < n; lnum++ {
		line, err = bf.ReadString('\n')
		if err == io.EOF {
			switch lnum {
			case 0:
				return "", errors.New("no lines in file")
			case 1:
				return "", errors.New("only 1 line")
			default:
				return "", fmt.Errorf("only %d lines", lnum)
			}
		}
		if err != nil {
			return "", err
		}
	}
	if line == "" {
		return "", fmt.Errorf("line %d empty", n)
	}
	return line, nil
}

func getMailInfo(fileName string) map[string]string {
	info := make(map[string]string)

	for i := 3; i <= 5; i++ {
		line, err := readSpecificLine(fileName, i)
		checkError(err)
		separatedLine := strings.Split(line, ":")

		if i == 5 && len(separatedLine) < 2 {
			break
		}

		if i == 5 && len(separatedLine) == 3 {
			separatedLine = append(separatedLine[:1], separatedLine[2:]...)
		}

		content := GetMailContent(fileName)

		key := strings.ToLower(separatedLine[0])
		value := separatedLine[1]

		info[key] = value
		info["content"] = content
	}
	return info
}

func GetMailContent(fileName string) string {
	f, err := os.Open(fileName)

	checkError(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var line int
	for scanner.Scan() {
		text := scanner.Text()
		line++
		if text == "" {
			break
		}
	}

	var content string

	for i := line + 1; scanner.Scan(); i++ {
		content += scanner.Text()
	}

	return content
}
