package intermediate

import (
	"fmt"
	"time"
)

func main() {

	// get current date time
	fmt.Println(time.Now())

	// specific time
	specificTime := time.Date(1999, 02, 26, 0, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)

	// Time parsing

	parsedTime, _ := time.Parse("2006-01-02", "1999-02-26")          // Mon Jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "99-02-26")             // Mon Jan 2 15:04:05 MST 2006
	parsedTime2, _ := time.Parse("06-1-2", "09-02-26")               // Mon Jan 2 15:04:05 MST 2006
	parsedTime3, _ := time.Parse("06-01-02 15-04", "99-02-26 18-09") // Mon Jan 2 15:04:05 MST 2006

	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)

	// Formatting time
	t := time.Now()
	fmt.Println("Formatted String:", t.Format("06-01-02 15:04:05 Mon"))

	// adding time
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	//Rounding off time
	t1 := time.Date(2024, time.July, 8, 14, 50, 40, 00, time.UTC)
	// fmt.Println("Rounded UTC Time", t1.Round(time.Hour))   // by Hour
	// fmt.Println("Rounded UTC Time", t1.Round(time.Minute)) // by Minute

	loc, _ := time.LoadLocation("Asia/Kolkata")

	//comnverting to specific time zone
	tLocal := t1.In(loc)

	roundTime := t1.Round(time.Hour)
	localRoundTime := tLocal.Round(time.Hour)

	fmt.Println("Local Time", tLocal)
	fmt.Println("UTC Time", t1)
	fmt.Println("rounded Local Time", localRoundTime)
	fmt.Println("rounded UTC Time", roundTime)

	//truncate time
	fmt.Println("Truncated Time:", t1.Truncate(time.Hour))

	loc1, _ := time.LoadLocation("America/New_York")

	//convert time to location

	tInNY := time.Now().In(loc1)

	fmt.Println("New York Time:", tInNY)

	//Duration between two date time
	time1 := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
	time2 := time.Date(2024, time.July, 4, 18, 0, 0, 0, time.UTC)

	duration := time2.Sub(time1)
	fmt.Println("Duration between two date time:", duration)

	//Compare time

	fmt.Println("Time2 is after time1 :", time2.After(time1))

}
