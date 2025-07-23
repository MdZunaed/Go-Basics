package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close() // This will close file after execution

	fileInfo, err := file.Stat() // Stat returns the [FileInfo] structure describing file

	if err != nil {
		panic(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Is directory:", fileInfo.IsDir())
	fmt.Println("File size:", fileInfo.Size())
	fmt.Println("Modified time:", fileInfo.ModTime())

	/// Read by buffer (Recommended)
	buff := make([]byte, fileInfo.Size())
	n, err := file.Read(buff) // n is length
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(buff); i++ {
		fmt.Println(n, string(buff[i])) // Just to check one by one byte
	}

	/// Read by built-in method (Not recommended for bigger file)
	data, err := os.ReadFile("test.txt")

	if err != nil {
		panic(err)
	}
	// no need to close
	fmt.Println("\n", string(data))

	/// Read directories
	dir, err := os.Open("./") // ../ will show previous directories
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	dirs, err := dir.ReadDir(0) // 0<= will return all
	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}

	/// Create File and write

	f, err := os.Create("create.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("Created with Go")
}
