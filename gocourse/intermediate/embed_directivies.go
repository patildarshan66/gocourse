package intermediate

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed example.txt
var exampleString string

//go:embed basics
var basicsFolder embed.FS

func main() {

	fmt.Println("exampleString:", exampleString)

	data, err := basicsFolder.ReadFile("basics/hello.go")
	if err != nil {
		fmt.Println("Error Reading a file:", err)
		return
	}
	fmt.Println("Content of basics/hello.go:\n", string(data))

	data1, err := basicsFolder.ReadDir("basics")
	if err != nil {
		fmt.Println("Error Reading a Dir:", err)
		return
	}
	for _, entry := range data1 {
		fmt.Println("Dir:", entry.Name(), entry.IsDir())
	}

	err = fs.WalkDir(basicsFolder, "basics", walkDirFunc)
	if err != nil {
		fmt.Println("Error Reading a Dir:", err)
		return
	}
}

func walkDirFunc(path string, d fs.DirEntry, err error) error {
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Path:", path)
	return nil
}
