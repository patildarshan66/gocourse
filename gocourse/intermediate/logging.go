package intermediate

import (
	"log"
	"os"
)

func main() {

	log.Println("This is a log message")
	log.SetPrefix("INFO: ")
	log.Println("This is a info log")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a set flag log")

	infoLogger.Println("This is a info logger")
	warningLogger.Println("This is a warning logger")
	errorLogger.Println("This is a error logger")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Failed to open the file: %v", err)
	}

	defer file.Close()

	infoLogger1 := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger1 := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	debuggingLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLogger1.Println("This is a info logger")
	warningLogger1.Println("This is a warning logger")
	errorLogger1.Println("This is a error logger")
	debuggingLogger.Println("This is a debug log")
}

var (
	infoLogger    = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger   = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
)
