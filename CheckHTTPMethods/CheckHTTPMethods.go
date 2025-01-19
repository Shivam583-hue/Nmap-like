package CheckHTTPMethods

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func CheckHTTPMethods(url string) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("HTTP Status: %s\n", resp.Status)
	fmt.Println("HTTP Headers:")
	for key, value := range resp.Header {
		fmt.Printf("%s: %s\n", key, strings.Join(value, ", "))
	}

	checkSupportedMethods(url)

	checkForDirectoryTraversal(url)
}

func checkSupportedMethods(url string) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "HEAD", "PATCH", "OPTIONS", "TRACE"}

	for _, method := range methods {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Method %s is allowed\n", method)
		}
	}
}

func checkForDirectoryTraversal(url string) {
	commonPaths := []string{
		"/admin", "/config", "/login", "/.git", "/wp-admin", "/uploads",
	}

	for _, path := range commonPaths {
		checkURL := fmt.Sprintf("%s%s", url, path)
		resp, err := http.Get(checkURL)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Potential accessible directory found: %s\n", checkURL)
		}
	}
}
