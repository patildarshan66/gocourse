package basics

import (
	"fmt"
	network "net/http"
)

func main() {
	resp, err := network.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response Status Code:", resp.StatusCode)
}
