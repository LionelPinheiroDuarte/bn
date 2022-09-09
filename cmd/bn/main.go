package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"time"
)

func main() {

	brain := filepath.Dir("/home/boost/projects/test/dump/")

	current_time := time.Now()
	var formatedDate = fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())

	fmt.Println(reflect.TypeOf(formatedDate))
	os.Chdir(brain)
	os.Mkdir(formatedDate, 0750)
	fmt.Println(formatedDate + "25")
	os.Chdir(formatedDate)
	fmt.Println(formatedDate + "27")
	_, err := os.Create("README.md")
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command("vi", "README.md")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
