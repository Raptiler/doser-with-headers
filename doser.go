package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	url             string
	payload         string
	threads         int
	requestCounter  int
	printedMsgs     []string
	waitGroup       sync.WaitGroup
	customHeaders   http.Header
)

func printMsg(msg string) {
	if !contains(printedMsgs, msg) {
		fmt.Printf("\n%s after %d requests\n", msg, requestCounter)
		printedMsgs = append(printedMsgs, msg)
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func init() {
	rand.Seed(time.Now().UnixNano())
	customHeaders = http.Header{}
}

func handleStatusCodes(statusCode int) {
	requestCounter++
	fmt.Printf("\r%d requests have been sent", requestCounter)

	if statusCode == 429 {
		printMsg("You have been throttled")
	}
	if statusCode == 500 {
		printMsg("Status code 500 received")
	}
}

func sendRequest(method string, client *http.Client, req *http.Request) {
	defer waitGroup.Done()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	handleStatusCodes(resp.StatusCode)
}

func makeRequest(method string) {
	client := &http.Client{}
	var req *http.Request
	var err error

	if method == "POST" {
		req, err = http.NewRequest(method, url, strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	for key, values := range customHeaders {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	sendRequest(method, client, req)
}

func main() {
	var headers string

	flag.StringVar(&url, "g", "", "Specify GET request. Usage: -g '<url>'")
	flag.StringVar(&url, "p", "", "Specify POST request. Usage: -p '<url>'")
	flag.StringVar(&payload, "d", "", "Specify data payload for POST request")
	flag.IntVar(&threads, "t", 500, "Specify number of threads to be used")
	flag.StringVar(&headers, "H", "", "Specify custom headers. Usage: -H 'Key: Value;Key2: Value2'")
	flag.Parse()

	if headers != "" {
		for _, header := range strings.Split(headers, ";") {
			parts := strings.SplitN(header, ":", 2)
			if len(parts) == 2 {
				customHeaders.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
			}
		}
	}

	if url == "" || (payload == "" && strings.Contains(url, "-p")) {
		fmt.Println("You must specify a URL with -g for GET or -p for POST and, if POST, include a payload with -d.")
		flag.Usage()
		return
	}

	waitGroup.Add(threads)

	for i := 0; i < threads; i++ {
		if payload != "" {
			go makeRequest("POST")
		} else {
			go makeRequest("GET")
		}
	}
	waitGroup.Wait()
}
