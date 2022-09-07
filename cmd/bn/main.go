package main

import (
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("Building my note command line tool")
	brain := filepath.Dir("~/projects/test/dump/")
	fmt.Println(brain + "/cap/")

	current_time := time.Now()
	var formatedDate string
	formatedDate = fmt.Sprintf("%d%02d%02d%02d%02d%02d\n",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())
	fmt.Println(formatedDate)
}
