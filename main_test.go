package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

// func TestService(t *testing.T) {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}

// 	url := os.Getenv("SERVICE_URL")
// 	if url == "" {
// 		url = "http://localhost:" + port
// 	}

// 	retryClient := retry.NewClient()
// 	req, err := retry.NewRequest(http.MethodGet, url+"/", nil)
// 	if err != nil {
// 		t.Fatalf("retry.NewRequest: %v", err)
// 	}

// 	token := os.Getenv("TOKEN")
// 	if token != "" {
// 		req.Header.Set("Authorization", "Bearer "+token)
// 	}

// 	resp, err := retryClient.Do(req)
// 	if err != nil {
// 		t.Fatalf("retryClient.Do: %v", err)
// 	}

// 	if got := resp.StatusCode; got != http.StatusOK {
// 		t.Errorf("HTTP Response: got %q, want %q", got, http.StatusOK)
// 	}

// 	out, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("io.ReadAll: %v", err)
// 	}

// 	fmt.Println("out: ", string(out))

// 	want := "Hello from Scout."
// 	if !strings.Contains(string(out), want) {
// 		t.Errorf("HTTP Response: body does not include %q", want)
// 	}
// }

func TestRapidAPI(t *testing.T) {
	url := "https://scout3.p.rapidapi.com/"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", "4581c59f6cmsh545a4f1797a2204p114ca1jsn6479f6982bd9")
	req.Header.Add("X-RapidAPI-Host", "scout3.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("http error: %v", err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		t.Errorf("HTTP Response: got %q, want %q", res.StatusCode, "200")
	}

	bodyStr := strings.Trim(string(body), " \n")

	if bodyStr != "Hello from Scout" {
		t.Errorf("HTTP Response: body does not include %q", "Hello from Scout")
	}
}
