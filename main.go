package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go_scaffold/copy"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatal("need at least one argument: project name")
		return
	}

	projectName := os.Args[1]
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	projectPath := filepath.Join(pwd, projectName)
	fmt.Println("生成项目路径：", projectPath)

	err = copy.CopyTo(projectPath, projectName)
	if err != nil {
		log.Fatal(err)
	}
}
