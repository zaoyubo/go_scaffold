package copy

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

const TEMPLATE_DIR = "../go_chassis_template"
const TMEPLATE_KEY_WORD = "template"

var box = packr.New("tmpl", TEMPLATE_DIR)

func CopyTo(toDir, projectName string) error {
	list := box.List()
	fmt.Println(list)
	for _, file := range list {
		toFileName := filepath.Join(toDir, file)
		err := ensureFileDirExist(toFileName)
		if err != nil {
			return err
		}
		content, err := getBoxFileContent(file)
		if err != nil {
			return err
		}
		err = CopyAndReplaceContent(toFileName, content, projectName)
		if err != nil {
			return err
		}
	}
	return nil
}

func ensureFileDirExist(file string) error {
	p := filepath.Dir(file)
	if _, e := os.Stat(p); e != nil {
		if e = os.MkdirAll(p, 0777); e != nil {
			return e
		}
	}
	return nil
}

func getBoxFileContent(filename string) (string, error) {
	content, err := box.FindString(filename)
	if err != nil {
		return "", err
	}
	return content, nil
}

func CopyAndReplaceContent(dst string, content string, projectName string) error {
	out, _ := os.Create(dst)
	defer out.Close()
	replaced := strings.Replace(content, TMEPLATE_KEY_WORD, projectName, -1)
	_, err := out.WriteString(replaced)
	if err != nil {
		return err
	}
	return nil
}
