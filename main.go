package main

import "fmt"

func main() {
	// Mebuat Var cara pertama
	var name string = "membuat variable pada golang"
	fmt.Println(name)

	// Mebuat Var cara kedua
	name1 := "Hai"
	fmt.Println(name1)

	// Mebuat Var cara ketiga
	name2 := 10
	fmt.Println(name2)

	// Variabel dengan tipe data dasar dan nilai eksplisit
	var bo bool = true          // Nilai: true
	var st string = "Hello"     // Nilai: "Hello"
	var in int = 42             // Nilai: 42
	var fl float64 = 3.14       // Nilai: 3.14
	var co complex128 = 1 + 2i  // Nilai: 1 + 2i
	var ru rune = 'A'           // Nilai: 'A' (Unicode)
	var byt byte = 255          // Nilai: 255

	// Menampilkan nilai-nilai variabel
	fmt.Println("Nilai variabel:")
	fmt.Println("Boolean:", bo)
	fmt.Println("String:", st)
	fmt.Println("Integer:", in)
	fmt.Println("Float64:", fl)
	fmt.Println("Complex128:", co)
	fmt.Println("Rune:", ru)
	fmt.Println("Byte:", byt)

	// Variabel dengan tipe data dasar dan nilai defaultnya
	var b bool         // Default: false
	var s string       // Default: ""
	var i int          // Default: 0
	var f float64      // Default: 0
	var c complex128   // Default: 0 + 0i
	var r rune         // Default: 0 (Unicode)
	var by byte        // Default: 0

	// Menampilkan nilai default
	fmt.Println("Nilai default:")
	fmt.Println("Boolean:", b)
	fmt.Println("String:", s)
	fmt.Println("Integer:", i)
	fmt.Println("Float64:", f)
	fmt.Println("Complex128:", c)
	fmt.Println("Rune:", r)
	fmt.Println("Byte:", by)

}