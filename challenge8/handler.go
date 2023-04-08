package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

func PostRequest() {
	data := Data{Water: h8HelperRand.RandomInt(1, 100), Wind: h8HelperRand.RandomInt(1, 100)}
	url := "https://jsonplaceholder.typicode.com/posts"

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err.Error())
	}

	buff := bytes.NewBuffer(requestJson)
	req, err := http.NewRequest("POST", url, buff)
	if err != nil {
		log.Fatalln(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer res.Body.Close()

	result := Data{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err.Error())
	}

	resJson, err := json.MarshalIndent(result, "", "")
	if err != nil {
		log.Fatalln(err.Error())
	}

	statusWater := WaterCheckLevel(data.Water)
	statusWind := WindCheckLevel(data.Wind)

	log.Println(string(resJson))
	fmt.Println("status water:", statusWater)
	fmt.Println("status wind:", statusWind)
	fmt.Println()
}
