package intermediate

import (
	"fmt"
	"net/url"
)

func main() {

	rawUrl := "https://darshanpatil.com:8080/path?query=param#fragment"

	parsedUrl, err := url.Parse(rawUrl)

	if err != nil {
		fmt.Println("Error Parsing URL:", err)
		return
	}

	fmt.Println("Parsed Url:", parsedUrl)
	fmt.Println("Parsed Url Scheme:", parsedUrl.Scheme)
	fmt.Println("Parsed Url Host:", parsedUrl.Host)
	fmt.Println("Parsed Url Port:", parsedUrl.Port())
	fmt.Println("Parsed Url Path:", parsedUrl.Path)
	fmt.Println("Parsed Url RawQuery:", parsedUrl.RawQuery)
	fmt.Println("Parsed Url Fragment:", parsedUrl.Fragment)

	rawUrl2 := "https://darshanpatil.com/path?name=Darshan&age=26"
	parsedUrl2, err := url.Parse(rawUrl2)

	if err != nil {
		fmt.Println("Error Parsing URL:", err)
		return
	}

	queryParams := parsedUrl2.Query()
	fmt.Println(queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	//Building URL

	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "darshanpatil.com",
		Path:   "/path",
	}

	query := baseUrl.Query()

	query.Set("name", "Darshan")
	query.Set("age", "26")
	baseUrl.RawQuery = query.Encode()

	fmt.Println("Built URL", baseUrl.String())

	value := url.Values{}

	value.Add("name", "Darshan")
	value.Add("age", "26")
	value.Add("city", "Navi Mumbai")
	value.Add("pincode", "700607")

	encodedQuery := value.Encode()

	fmt.Println("Encoded Query:", encodedQuery)

	newUrl := "https://darshanpatil.com/search"
	fullUrl := newUrl + "?" + encodedQuery

	fmt.Println("Full Url:", fullUrl)
}
