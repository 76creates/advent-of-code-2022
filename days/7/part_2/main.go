package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	dirs        = make(map[string]map[string]int)
	currentDir  []string
	totalFsSize = 70000000
	needUnused  = 30000000
)

func main() {
	f, _ := os.Open("../input")
	defer f.Close()

	s := bufio.NewScanner(f)
	var output []string
	for s.Scan() {
		line := s.Text()
		output = append(output, line)
	}

	dirs["/"] = make(map[string]int)
	for {
		if len(output) == 0 {
			break
		}
		out := output[0]
		output = output[1:]

		if strings.HasPrefix(out, "$ cd") {
			toDir := strings.TrimPrefix(out, "$ cd ")
			if toDir == "/" {
				currentDir = []string{}
			} else if toDir == ".." {
				currentDir = currentDir[:len(currentDir)-1]
			} else {
				currentDir = append(currentDir, toDir)
			}
		} else if out == "$ ls" {
			var lsData []string
			for {
				if len(output) == 0 {
					break
				}
				lsLine := output[0]
				if strings.HasPrefix(lsLine, "$") {
					break
				}
				lsData = append(lsData, lsLine)
				output = output[1:]
			}
			for _, lsLine := range lsData {
				if strings.HasPrefix(lsLine, "dir") {
					dirName := strings.TrimPrefix(lsLine, "dir ")
					dirs[getCurrentDir()+dirName+"/"] = make(map[string]int)
					dirs[getCurrentDir()][dirName] = -1
				} else {
					f := strings.Split(lsLine, " ")
					fileName := f[1]
					fileSize, _ := strconv.Atoi(f[0])
					dirs[getCurrentDir()][fileName] = fileSize
				}
			}
		} else {
			panic(fmt.Sprintf("unknown command: %s", out))
		}
	}

	total := getDirSize("/")
	neededSpace := needUnused - (totalFsSize - total)
	fmt.Printf("fs: %d, used: %d, needed: %d, need to free: %d\n", totalFsSize, total, needUnused, neededSpace)
	var perfectDirForDeleting string
	var perfectDirSize int
	for k, _ := range dirs {
		if size := getDirSize(k); size >= neededSpace {
			if size == neededSpace {
				perfectDirForDeleting = k
				break
			} else {
				if perfectDirSize == 0 || size < perfectDirSize {
					perfectDirForDeleting = k
					perfectDirSize = size
				}
			}
		}
	}
	fmt.Println("perfect dir for deleting:", perfectDirForDeleting)
	fmt.Println("size of perfect dir:", perfectDirSize)
}

func getCurrentDir() string {
	if len(currentDir) == 0 {
		return "/"
	}
	return "/" + strings.Join(currentDir, "/") + "/"
}

func getDirSize(dir string) int {
	size := 0
	for k, v := range dirs[dir] {
		if v == -1 {
			size += getDirSize(dir + k + "/")
		} else {
			size += v
		}
	}
	return size
}
