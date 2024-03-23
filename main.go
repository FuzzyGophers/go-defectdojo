package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

const defaultBaseURL string = "http://127.0.0.1:8080/"
const apiPath string = "api/v2/"
const token = "c45e3e9b429aed99f85165fdf05cd5ddadd93753"
const finalToken string = "Token " + token
const contentType string = "application/json"

func main() {

	req, err := http.NewRequest(http.MethodGet, defaultBaseURL+apiPath+"product_types/", nil)

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", finalToken)

	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: response body: %s\n", resBody)
}

type Client struct {
	client *retryablehttp.Client
	method string
	path   string
	token  string
}
