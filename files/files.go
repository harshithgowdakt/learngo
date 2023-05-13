package files

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) {
	file, err :=os.Open(path);
	if err!=nil {
		fmt.Println("Error opening file" + err.Error());
	}

	defer file.Close();

	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanLines);


	for scanner.Scan() {
		line := scanner.Text();
		fmt.Println(line);
	}

	if err:=scanner.Err(); err!=nil{
		fmt.Println("Error while scaning file data" + err.Error());
	}
}