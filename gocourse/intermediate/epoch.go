package intermediate

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	// convert time to epoch
	epochT := now.Unix()
	fmt.Println("Current Epoch Time:", epochT)

	// convert epoch to time
	t := time.Unix(epochT, 0)
	fmt.Println("converted Time:", t)

}
