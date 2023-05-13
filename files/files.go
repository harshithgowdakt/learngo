package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file" + err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while scaning file data" + err.Error())
	}
}

func ReadFileUsingIoUtil(path string) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Error while reading file" + err.Error())
		return;
	}

	fmt.Println(string(data));
}
