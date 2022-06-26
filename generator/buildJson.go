package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func buildJson(jsonNumbers int) {
	log.Printf("Start to generate %d json files\n", jsonNumbers)

	for i := 1; i <= jsonNumbers; i++ {
		fileName := fmt.Sprintf("json/%d.json", i)
		_, err := os.Stat(fileName)
		// if the file already exists, don't generate
		if !os.IsNotExist(err) {
			log.Printf("%v file exists. ❌ \n", fileName)
			continue
		}
		// force random with time
		rand.Seed(time.Now().UnixNano())
		content := fmt.Sprintf(`{"linked_schemas": ["test_schema-v1"],"geolocation": {"lat": %d,"lon": %d},"name": "Test-%d"}`, rand.Intn(90), rand.Intn(180), i)
		write(fileName, content)
		log.Printf("%v is generated. ✅ \n", fileName)
	}

	log.Printf("Congratulations! 🎉 Successfully generate %d json files\n", jsonNumbers)
}

func write(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
