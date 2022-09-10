package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func formatedDate() string {
	current_time := time.Now()
	var formatedDate = fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())
	return formatedDate
}
func createDirectory(title string) {
	os.Mkdir(title, 0750)
	fmt.Println(title)
	os.Chdir(title)
	fmt.Println(title)
}
func main() {

	brain := filepath.Dir("/home/boost/projects/test/dump/")

	formatedDate := formatedDate()

	os.Chdir(brain)
	createDirectory(formatedDate)

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)

	file, err := os.Create("README.md")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	lscmd := exec.Command("ls")
	output, err := lscmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Output \n%s\n", string(output))

	fmt.Println("exe vi")
	cmd := exec.Command("vim", "README.md")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	oops := cmd.Run()
	if err != nil {
		log.Fatal(oops)
	}
}
