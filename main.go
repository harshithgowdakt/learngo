package main

import (
	"fmt"
	"os"
	"path"

	"github.com/harshithgowdakt/learngo/files"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    cwd, err := os.Getwd();
    if err!=nil{
        fmt.Println("Error while getting current working directory" + err.Error());
    }

    path := path.Join(cwd, "files/data");
    fmt.Println("File path is : " + path);

    //read file
    files.ReadFile(path);
}