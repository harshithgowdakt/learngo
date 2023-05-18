package main

import (
	"fmt"
	"os"
	"path"

	"github.com/harshithgowdakt/learngo/ds"
	"github.com/harshithgowdakt/learngo/files"
	"github.com/harshithgowdakt/learngo/oops"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory" + err.Error())
	}

	filepath := path.Join(cwd, "files/data")
	fmt.Println(" File filepath is : " + filepath)

	fmt.Println("####################### Read file using bufio and os ############################## ")

	//read file
	files.ReadFile(filepath)

	fmt.Println("####################### Read file using ioutil ############################## ")

	files.ReadFileUsingIoUtil(filepath)

	fmt.Println("####################### Read file using os readfile ###########################")
	files.ReadFileUingOs(filepath)


	// create folder resource if not exists
	err = os.MkdirAll(path.Join(cwd, "resource"), os.ModePerm)

	if err!=nil {
		fmt.Println("Error while creating directory : ", err.Error())
	}

	fmt.Println("####################### write to file ###########################")
	files.WriteUsingIoUtil(path.Join(cwd, "resource/write_ioutil"))

	fmt.Println("####################### write to file ###########################")
	files.WriteUsingOs(path.Join(cwd, "resource/write_os"))

	fmt.Println("####################### write to file ###########################")
	files.WriteUsingBuf(path.Join(cwd, "resource/write_buf"))


	fmt.Println("Running oops example")
	oops.RunStoreExample()

	fmt.Println("running slice examples")
	ds.RunSliceExmaples()

	fmt.Println("Running map example")
	ds.RunMapExample()
}