package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("hello.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fileinfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file", fileinfo.Name())
	// fmt.Println("Is folder", fileinfo.Size())
	// fmt.Println("Updation time: ", fileinfo.ModTime())
	// fmt.Println("Permission", fileinfo.Mode())

	//------------One way to read file
	// buf := make([]byte, 100)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	fmt.Println("data", d, string(buf[i]))
	// }

	// if err != nil {
	// 	panic(err)
	// // }

	// defer f.Close()

	//--------------Another way to read the file
	// data, err := os.ReadFile("hello.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	//------------checking the files in the directory
	// fileinfo, err := dir.ReadDir(-1)

	// for _, fi := range fileinfo {
	// 	fmt.Println(fi.Name())
	// }

	//Creating a file and writing in the file and reading it
	// f, err := os.Create("welcome.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hi, go!")
	// f.WriteString("I am leaving from Go")

	// data, err := os.ReadFile("welcome.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	//-----------------read & write to another fule using stream fashion
	// sourceFile, err := os.Open("hello.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()
	// destfile, err := os.Create("yo.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destfile)
	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	writer.WriteByte(b)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("Written to file")

	//------------Deleting the file
	err := os.Remove("yo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file deleted")

}
