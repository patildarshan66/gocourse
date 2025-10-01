package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader:", err)
	}
	fmt.Println("Read bytes:", string(buf[:n]))

}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error reading from reader:", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error closing resource:", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer
	data := "Hello, Buffer!"
	_, err := buf.WriteString(data)
	if err != nil {
		log.Fatalln("Error writing to buffer:", err)
	}
	fmt.Println("Wrote bytes:", buf.String())
}

func multipleReadersExample() {
	r1 := strings.NewReader("Hello, ")
	r2 := strings.NewReader("World!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	if err != nil && err != io.EOF {
		log.Fatalln("Error reading from multi-reader:", err)
	}
	fmt.Println("Read bytes:", buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello, Darshan"))
		pw.Close()
	}()
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(pr)
	if err != nil {
		log.Fatalln("Error reading from pipe:", err)
	}
	fmt.Println("Read bytes from pipe:", buf.String())
}

func writeToFileExample(filepath string, data string) {

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to file:", err)
	}
	fmt.Println("Wrote data to file:", filepath)
	myResource := &MyResource{name: "MyCustomResource"}
	closeResource(myResource)

	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln("Error writing to file via Writer interface:", err)
	// }
	// fmt.Println("Wrote data to file via Writer interface:", filepath)

}

type MyResource struct {
	name string
}

func (r *MyResource) Close() error {
	fmt.Println("Closing resource:", r.name)
	return nil
}

func main() {
	fmt.Println("Read from Reader example:")
	readFromReader(strings.NewReader("Hello, Reader!"))

	fmt.Println("\nWrite to Writer example:")
	var buf bytes.Buffer
	writeToWriter(&buf, "Hello, Writer!")
	fmt.Println("Wrote bytes:", buf.String())

	fmt.Println("\nBuffer example:")
	bufferExample()

	fmt.Println("\nMultiple Readers example:")
	multipleReadersExample()

	fmt.Println("\nPipe example:")
	pipeExample()

	fmt.Println("\nWrite to File example:")
	writeToFileExample("io.txt", "Hello, File!\n")
}
