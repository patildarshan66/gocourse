package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}
func main() {
	checkError(os.Mkdir("advanced", 0755))

	checkError(os.WriteFile("advanced/file.txt", []byte("Hello, Darshan"), 0755))
	checkError(os.MkdirAll("subdir/parent/child", 0755))
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))

	checkError(os.WriteFile("subdir/parent/child1/file.txt", []byte("Hello, Darshan"), 0755))
	checkError(os.WriteFile("subdir/parent/file.txt", []byte("Hello, Darshan"), 0755))

	result, err := os.ReadDir("subdir/parent")

	checkError(err)

	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir())
	}

	checkError(os.Chdir("subdir/parent/child"))
	dir, err := os.Getwd()

	checkError(err)
	fmt.Println("Current Working Directory:", dir)

	result, err = os.ReadDir(".")

	checkError(err)

	fmt.Println("Reading subdir/parent/child")
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir())
	}

	checkError(os.Chdir("../../../"))
	dir, err = os.Getwd()

	checkError(err)
	fmt.Println("Current Working Directory:", dir)

	//waking through directory
	fmt.Println("waking through directory")
	filePath := "subdir"
	err = filepath.WalkDir(filePath, walkDirFunc)
	checkError(err)

	//waking
	fmt.Println("waking")
	err = filepath.Walk(filePath, walkFunc)
	checkError(err)

	checkError(os.RemoveAll("subdir")) //to delete directory and it's all files

	checkError(os.Remove("advanced/file.txt")) //to delete specific file or empty directory

	checkError(os.Remove("advanced")) //to delete specific file or empty directory

}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println(path)
	return nil
}
func walkDirFunc(path string, d os.DirEntry, err error) error {
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println(path)
	return nil
}
