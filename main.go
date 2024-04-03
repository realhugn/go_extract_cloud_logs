package main

import (
	config "data_extract/database"
	"data_extract/model"
	"data_extract/utils"
	"fmt"
	"log"
	"path/filepath"
	"sync"
)

func main() {
	db := config.ConnectDB()
	db.AutoMigrate(&model.Infomation{})
	matches, err := filepath.Glob("/home/stackops/data/2024-04-01_a_@logsinspector/*/Passwords.txt")
	matches2, _ := filepath.Glob("/home/stackops/data/2024-04-01_a_@logsinspector/*/passwords.txt")
	var wg sync.WaitGroup
	if err != nil {
		log.Fatal(err)
	}
	for _, filename := range matches {
		wg.Add(1)
		go func(filename string) {
			content := utils.Read_file(filename)
			utils.Extract_info(content, db, "20240402")
		}(filename)
	}
	wg.Wait()
	for _, filename := range matches2 {
		wg.Add(1)
		go func(filename string) {
			content := utils.Read_file(filename)
			utils.Extract_info(content, db, "20240402")
		}(filename)
	}
	wg.Wait()
	fmt.Println(len(matches))

}
