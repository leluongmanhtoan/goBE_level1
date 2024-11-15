package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Data struct {
	Id         int64  `json:"id"`
	Name       string `json:"employee_name"`
	Salary     int64  `json:"employee_salary"`
	Age        int    `json:"employee_age"`
	ProfileImg string `json:"profile_image"`
}

type employee struct {
	Status string `json:"status"`
	Data   []Data `json:"data"`
	Mess   string `json:"message"`
}

func main() {
	url := "https://dummy.restapiexample.com/api/v1/employees"
	cookie := &http.Cookie{
		Name:     "humans_21909",
		Value:    "1",
		Path:     "/api/v1",
		Expires:  time.Date(2025, 11, 15, 5, 29, 2, 0, time.UTC),
		HttpOnly: true,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("create request failed: %v", err)
	}
	req.AddCookie(cookie)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("call API failed: %v", err)
	}
	if res == nil {
		log.Fatalf("response empty!")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("read response failed: %v", err)
	}

	if contentType := res.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
		log.Fatalf("unexpect content type")
	}

	//fmt.Println("Response Body:", string(body))
	var employeeResponse employee
	err = json.Unmarshal(body, &employeeResponse)

	if err != nil {
		log.Fatalf("parse JSON failed: %v", err)
	}
	fmt.Println(employeeResponse)
}
