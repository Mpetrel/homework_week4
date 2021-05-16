package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"homework_week4/internal/conf"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	execPath, err := os.Executable()
	fmt.Println(execPath)
	str := filepath.Join(filepath.Dir(execPath), "../../configs/config.yaml")
	fmt.Println(str)
	yamlFile, err := ioutil.ReadFile("E:\\dev\\go_workspace\\homework_week4\\configs\\config.yaml")
	if err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err = yaml.Unmarshal(yamlFile, &bc); err != nil {
		panic(err)
	}

	fmt.Println(bc.Data)
}
