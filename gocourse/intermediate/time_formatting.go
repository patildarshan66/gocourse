package intermediate

import (
	"fmt"
	"time"
)

func main() {

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2025-07-26T23:30:00Z"

	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println("Error while parsing date:", err)
		return
	}
	fmt.Println(t)

	str1 := "Jul 03 2025 12:05 PM"
	layout2 := "Jan 02 2006 15:04 PM"
	t1, err := time.Parse(layout2, str1)
	fmt.Println(t1)
}
