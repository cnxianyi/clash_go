package config

import (
	"bufio"
	"clash_go/tp"
	"fmt"
	"io"
	"os"
	"strings"
)

func Setup() {
	// 打开文件
	configFile, err := os.Open("./config/config.yaml")
	defer configFile.Close()
	if err != nil {
		fmt.Println("读取config.yaml 失败")
	}

	// 创建bufio reader 进行 每行 读取
	reader := bufio.NewReader(configFile)

	// 创建容纳512行的 字符串切片
	// l := make([]string, 0, 512)

	var fa string

	var basic tp.Clash

	for {
		// 处理每一行
		line, _, err := reader.ReadLine()
		if err != nil {
			// 如果读取完毕则退出
			if err == io.EOF {
				break
			} else {
				// 读取异常
				fmt.Println(err)
			}
		}

		// 如果 第一个字符为空格
		if string(line)[0] == 32 { // 换行
			// 清除左侧空格
			res := strings.TrimLeft(string(line), " ")
			fmt.Println(fa, ":", res)

		} else {
			// 根据 : 进行分割字符串
			before, after, found := strings.Cut(string(line), ":")
			if found {
				// : 右侧为空格或空
				if after == " " || after == "" {
					fa = before
				} else {
					basic = tp.Tag(&basic, before, after)
				}
			} else {
				fmt.Println("格式错误", string(line))
			}
		}

		// l = append(l, strings.Trim(string(line), "\n"))
	}

	fmt.Println(basic)
}
