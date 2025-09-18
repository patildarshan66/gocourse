package intermediate

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()

	// Set Log Level
	log.SetLevel(logrus.InfoLevel)

	// Set Log Format
	log.SetFormatter(&logrus.JSONFormatter{})

	//Logging examples
	log.Info("This is a info message.")
	log.Warn("This is a warning message.")
	log.Error("This is a error message")

	// Logging with fields
	log.WithFields(logrus.Fields{"username": "Darshan", "method": "GET"}).Info("User Logged in.")

}
