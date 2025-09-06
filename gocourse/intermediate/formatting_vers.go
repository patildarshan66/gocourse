package intermediate

import "fmt"

func main() {

	// ---General Formatting---
	// %v - the value in a default format
	// %+v - the value in a default format with field names
	// %#v - a Go-syntax representation of the value
	// %T - a Go-syntax representation of the type of the value
	// %% - a literal percent sign; consumes no value

	i := 12_345.67
	nameD := "Darshan"

	fmt.Println("General Formatting:")
	fmt.Printf("i: %v\n", i)   // Default format
	fmt.Printf("i: %#v\n", i)  // Go-syntax representation
	fmt.Printf("i: %T\n", i)   // Type of the value
	fmt.Printf("i: %v%%\n", i) // Literal percent sign

	fmt.Printf("name: %v\n", nameD)
	fmt.Printf("name: %#v\n", nameD)
	fmt.Printf("name: %T\n", nameD)

	// -- Integer Formatting---
	// %b - base 2
	// %d - base 10
	// %d+ - base 10 with sign
	// %o - base 8
	// %O - base 8 with leading 0o
	// %x - base 16, with lower-case letters for a-f
	// %X - base 16, with upper-case letters for A-F
	// %#x - base 16, with leading 0x
	// %4d - pad with spaces to width 4 , right-justified
	// %-4d - pad with spaces to width 4, left-justified
	// %04d - pad with leading zeros to width 4

	n := 42

	fmt.Println("\nInteger Formatting:")
	fmt.Printf("n: %b\n", n)   // Base 2
	fmt.Printf("n: %d\n", n)   // Base 10
	fmt.Printf("n: %+d\n", n)  // Base 10 with sign
	fmt.Printf("n: %o\n", n)   // Base 8
	fmt.Printf("n: %O\n", n)   // Base 8 with leading 0o
	fmt.Printf("n: %x\n", n)   // Base 16, lower-case
	fmt.Printf("n: %X\n", n)   // Base 16, upper-case
	fmt.Printf("n: %#x\n", n)  // Base 16 with leading 0x
	fmt.Printf("n: %4d\n", n)  // Pad with spaces to width 4, right-justified
	fmt.Printf("n: %-4d\n", n) // Pad with spaces to width 4, left-justified
	fmt.Printf("n: %04d\n", n) // Pad with leading zeros to width 4

	// --- String Formatting ---

	name := "Darshan"
	fmt.Printf("%s\n", name)   // Default string format
	fmt.Printf("%q\n", name)   // Double-quoted string
	fmt.Printf("%#v\n", name)  // Go-syntax representation
	fmt.Printf("%T\n", name)   // Type of the value
	fmt.Printf("%8s\n", name)  // Right-justified in a field of width 8
	fmt.Printf("%-8s\n", name) // Left-justified in a field of width 8
	fmt.Printf("%x\n", name)   // Hexadecimal representation
	fmt.Printf("% x\n", name)  // Hexadecimal with spaces between bytes

	// ---boolean Formatting---
	b := true
	fmt.Printf("%t\n", b)  // Default format
	fmt.Printf("%v\n", b)  // Go-syntax representation
	fmt.Printf("%T\n", b)  // Type of the value
	fmt.Printf("%#v\n", b) // Go-syntax representation

	// ---Floating-Point and Complex Number Formatting---
	f := 918.00
	fmt.Printf("%f\n", f)    // Decimal point, no exponent, default precision
	fmt.Printf("%e\n", f)    // Scientific notation, e.g., -1.234456e+78
	fmt.Printf("%g\n", f)    // Exponent as needed, only necessary digits
	fmt.Printf("%.2f\n", f)  // Decimal point, no exponent, default precision
	fmt.Printf("%6.2f\n", f) // Width 6, precision 2

}
