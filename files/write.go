package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteUsingIoUtil(path string) {
	data := []byte("Hello world")

	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println("Error while writing file" + err.Error())
	}
}

func WriteUsingOs(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error while creating file : ", err)
	}
	data := "Testing write operation"

	_, err = file.Write([]byte(data))
	if err != nil {
		fmt.Println("Error while writing to file", err)
	}
}

func WriteUsingBuf(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error while creating file : ", err)
		return
	}

	writer := bufio.NewWriter(file)
	data := "Testing writing operation using bufio"

	writer.Write([]byte(data))
	n, err := writer.Write([]byte(data))
	if err != nil {
		fmt.Println("Error while writing to file : ", err)
		return
	}

	fmt.Println("Successfully wrote, number of bytes : ", n)
}
