package FindIp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FindIp() {
	url := "https://api.ipify.org?format=json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(body))
}
