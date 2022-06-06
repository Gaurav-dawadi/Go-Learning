package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	////// Check if directory to store all files exists. If not, first create them.
	if _, err := os.Stat("media"); os.IsNotExist(err) {
		if err := os.Mkdir("media", 0777); err != nil {
			fmt.Println("Error occured during creation of directory: ", err)
		}
	}

	dirInfo, _ := os.Stat("media")
	// fmt.Println(dirInfo)

	fileExtension := filepath.Ext("myImagefile.txt")
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10) + fileExtension

	fileInfo, err := os.Create(filepath.Join(dirInfo.Name(), fileName))
	if err != nil {
		fmt.Println("Error occured during creation of file: ", err)
	}

	// file_slice := strings.Split(fileInfo.Name(), "/")
	// for idx, v := range file_slice {
	// 	fmt.Println(idx, v)
	// }
	// fmt.Println(file_slice[1])

	file_open, err := os.OpenFile(fileInfo.Name(), os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("An Error has occured during opening of file")
	}
	defer file_open.Close()

	text := "Golang offers a vast inbuilt library that can be used to perform read and write operations on files. In order to read from files on the local system, the io/ioutil module is put to use. The io/ioutil module is also used to write content to the file. "
	// // text := "\n" + "The fmt module implements formatted I/O with functions to read input from the stdin and print output to the stdout. The log module implements simple logging package. It defines a type, Logger, with methods for formatting the output. The os module provides the ability to access native operating-system features. The bufio module implements buffered I/O which helps to improve the CPU performance."

	write_file, err := file_open.WriteString(text)
	if err != nil {
		fmt.Println("An Error has occured during writing of file")
	}
	fmt.Println("Write file: ", write_file)

	data, err := ioutil.ReadFile(fileInfo.Name())
	if err != nil {
		fmt.Println("An Error has occured during writing of file")
	}
	fmt.Printf("%s\n", data)
}
