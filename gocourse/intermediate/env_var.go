package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var:", user)
	fmt.Println("Home env var:", home)

	err := os.Setenv("FRUIT", "MANGO")

	if err != nil {
		fmt.Println("Error Setting env var:", err)
	}
	fmt.Println("Set FRUIT Env Var successfully")

	fmt.Println("FRUIT env ver:", os.Getenv("FRUIT"))

	for _, e := range os.Environ() {
		kvPair := strings.SplitN(e, "=", 2)
		fmt.Println(kvPair[0])
	}

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting env var:", err)
		return
	}
	fmt.Println("Unset FRUIT Env Var successfully")
	fmt.Println("FRUIT env ver:", os.Getenv("FRUIT"))

	str := "a=b=c=d=e"

	// n = 1 returns "a=b=c=d=e"
	// n = 2 returns "a" and "b=c=d=e"
	// n = 3 returns "a" and "b" and "c=d=e"
	// n = 4 returns "a" and "b" and "c" and "d=e"
	// n = 5 returns "a" and "b" and "c" and "d" and "e"

	fmt.Println(strings.SplitN(str, "=", -1))
	fmt.Println(strings.SplitN(str, "=", 0))
	fmt.Println(strings.SplitN(str, "=", 1))
	fmt.Println(strings.SplitN(str, "=", 2))
	fmt.Println(strings.SplitN(str, "=", 3))
	fmt.Println(strings.SplitN(str, "=", 4))
	fmt.Println(strings.SplitN(str, "=", 5))

}
