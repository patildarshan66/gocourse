package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~lo, Good Morning!")

	// Encode to base64
	encodedStr := base64.StdEncoding.EncodeToString(data)

	fmt.Println("Encoded base64:", encodedStr)

	// Decode base64 to slice (bytes)

	decodedData, err := base64.StdEncoding.DecodeString(encodedStr)

	if err != nil {
		fmt.Println("Error in decoding base64:", err)
		return
	}

	fmt.Println("Decoded Data:", decodedData)
	fmt.Println("Decoded String:", string(decodedData))

	/// URL safe, avoid '/' and '+' while encoding a string

	encodedStr1 := base64.URLEncoding.EncodeToString(data)

	fmt.Println("URL safe Encoded base64:", encodedStr1)

}
