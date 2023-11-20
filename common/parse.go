package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadTsp(path string) []Node {
	nodes := make([]Node, 0)
	file, err := os.Open(path)
	if err != nil {
		panic("failed to read file")
	}
	defer file.Close()
	// 创建一个 Scanner 对象，用于逐行读取文件
	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	fmt.Println("reading file...")
	for i := 0; i < 5; i ++ {
		fmt.Println(lines[i])
	}
	for i := 6; i < len(lines) - 1; i ++ {
		nodeInfo := strings.Split(lines[i], " ")
		id, _ := strconv.Atoi(nodeInfo[0])
		x, _ := strconv.ParseFloat(nodeInfo[1], 64)
		y, _ := strconv.ParseFloat(nodeInfo[2], 64)
		// id从0开始
		nodes = append(nodes, Node{Id: id - 1, X: x, Y: y})
	}
	return nodes
}