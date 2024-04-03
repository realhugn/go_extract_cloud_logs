package utils

import (
	"bytes"
	"data_extract/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"gorm.io/gorm"
)

func Read_file(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func Extract_info(content []byte, db *gorm.DB, key string) {
	re := regexp.MustCompile(`
URL: (.*)
Username: (.*)
Password: (.*)
Application: (.*)`)
	content_string := string(content[:])

	matches := re.FindAllStringSubmatch(content_string, -1)

	for _, match := range matches {
		db.Table("infomations").Create(&model.Infomation{
			Url:         match[1],
			Username:    match[2],
			Password:    match[3],
			Application: match[4],
			Partkey:     key,
		})
		// if rs.Error == nil {
		// 	fmt.Println("Inserted")
		// }

	}
	fmt.Println("Inserted ", len(matches))
}

func MakeRequest(url string, password string, username string) {
	postBody, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}
