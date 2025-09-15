package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	pass := "password1234"

	// hash256 := sha256.Sum256([]byte(pass))

	// hash512 := sha512.Sum512([]byte(pass))

	// fmt.Println("Hash256:", hash256)
	// fmt.Printf("Hash256 Hex: %x\n", hash256)

	// fmt.Println("Hash512:", hash512)
	// fmt.Printf("Hash512 Hex: %x\n", hash512)

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error in generating salt:", err)
		return
	}
	fmt.Printf("Orignal Salt: %x\n", salt)

	//hash the password
	hashPass := hashPassword(salt, pass)

	//store the salt and password in databasem, just printing as of now

	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)
	fmt.Println("Hash:", hashPass)

	//verify pass
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)

	if err != nil {
		fmt.Println("Error while decoding Salt:", err)
		return
	}
	fmt.Printf("DecodedSalt Orignal Salt: %x\n", decodedSalt)
	pass1 := "password12345"
	loginPassHash := hashPassword(decodedSalt, pass1)

	if loginPassHash == hashPass {
		fmt.Println("Password matched successfully")
	} else {
		fmt.Println("Password incorrect!")
	}

}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)

	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(salt []byte, pass string) string {
	saltedPass := append(salt, []byte(pass)...)
	hash := sha256.Sum256(saltedPass)
	return base64.StdEncoding.EncodeToString(hash[:])
}
