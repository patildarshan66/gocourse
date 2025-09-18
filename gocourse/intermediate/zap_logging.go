package intermediate

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()

	if err != nil {
		fmt.Println("Error in initilizing logger:", err)
	}

	defer logger.Sync()

	logger.Info("This is an Info message")
	logger.Info("User Logged In", zap.String("username", "Darshan"), zap.String("method", "GET"))
}
