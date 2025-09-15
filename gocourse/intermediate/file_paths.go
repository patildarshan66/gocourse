package intermediate

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

	joinedPath := filepath.Join("home", "user", "download", "golang", "file.go")
	fmt.Println("joinedPath:", joinedPath)

	cleanedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("cleanedPath:", cleanedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")

	fmt.Println("Dir:", dir)
	fmt.Println("file:", file)
	fmt.Println("Base Path:", filepath.Base("/home/user/docs/file.txt"))
	fmt.Println("Base Path:", filepath.Base("/home/user/docs/"))

	fmt.Println("relativePath is absolute:", filepath.IsAbs(relativePath))
	fmt.Println("absolutePath is absolute:", filepath.IsAbs(absolutePath))

	fmt.Println("File Extension:", filepath.Ext("/home/user/docs/file.txt")) // to get file extension

	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file))) // to get file name without extension

	rel, err := filepath.Rel("/home/docs", "/home/downloads/file.txt") // return relative path to access `/home/downloads/file.txt` from `/home/docs`

	if err != nil {
		panic(err)
	}
	fmt.Println("Relative Path:", rel)

	rel, err = filepath.Rel("/home/docs", "/home/docs/go/file.txt") // remove path till same path like in this case `/home/docs/` is same in both

	if err != nil {
		panic(err)
	}
	fmt.Println("Relative Path:", rel)

	absPath, err := filepath.Abs(relativePath) // convert relative path to relative path

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Absolute Path:", absPath)

}
