package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func postJson(jsonNumbers int) {
	log.Printf("Start to post %d json files\n", jsonNumbers)

	indexUrl := viperEnvVariable("INDEX_URL")
	githubPageUrl := viperEnvVariable("GITHUB_PAGE_URL")

	for i := 1; i <= jsonNumbers; i++ {
		fileName := fmt.Sprintf("json/%d.json", i)

		body, _ := json.Marshal(map[string]string{
			"profile_url": githubPageUrl + "/" + fileName,
		})
		reqBody := bytes.NewBuffer(body)
		res, err := http.Post(indexUrl, "application/json", reqBody)
		if err != nil {
			log.Fatalf("Post file %s failed: %v. ❌", fileName, err)
		}
		defer res.Body.Close()
		if res.StatusCode == 200 {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatalln(err)
			}
			bodyString := string(body)
			log.Printf("%v response: %v\n", fileName, bodyString)
			log.Printf("%v/%v is posted. ✅ \n", githubPageUrl, fileName)
		}
	}

	log.Printf("Congratulations! 🎉 Successfully post %d json files\n", jsonNumbers)
}
