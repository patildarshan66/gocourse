package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("Error opening a file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	keyword := "important"

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedKeyword := strings.ReplaceAll(line, keyword, "neccessary")
			fmt.Println("Filtered Line:", line)
			fmt.Println("updated keyword Line:", updatedKeyword)
		}
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Error reading a file:", err)
		return
	}

}
