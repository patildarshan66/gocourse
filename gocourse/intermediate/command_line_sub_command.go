package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	stringFlag := flag.String("user", "Guest", "Name of the User")
	flag.Parse()
	fmt.Println(*stringFlag)

	subCommand1 := flag.NewFlagSet("firstsub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondsub", flag.ExitOnError)

	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1023, "Byte length of result")

	language := subCommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstsub":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		fmt.Println("processing:", *firstFlag)
		fmt.Println("Bytes:", *secondFlag)
	case "secondsub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("Language:", *language)
	default:
		fmt.Println("No Subcommand Entered")
		os.Exit(1)
	}
}
