package oops

import "fmt"

type store interface {
	read(duid string)
	write(data string, uid string)
	update(data string, uid string)
	delete(uid string)
}

type filestore struct {
}

func NewFileStore() *filestore {
	return &filestore{}
}

func (f *filestore) write(data string, uid string) {
	// will write implementation later
	fmt.Println("Writing data to file")
}

func (f *filestore) update(data string, uid string) {
	// will write implementation later
	fmt.Println("Updating data in the file")
}

func (f *filestore) read(uid string) {
	// will write implementation later
	fmt.Println("reading data from the file")
}

func (f *filestore) delete(uid string) {
	// will write implementation later
	fmt.Println("deleting data from the file")
}

type s3store struct {
}

func NewS3Store() *s3store {
	return &s3store{}
}

func (s *s3store) write(data string, uid string) {
	// will write implementation later
	fmt.Println("Writing data to s3")
}

func (s *s3store) update(data string, uid string) {
	// will write implementation later
	fmt.Println("updating data in the s3")
}

func (s *s3store) read(uid string) {
	// will write implementation later
	fmt.Println("reading data from s3")
}

func (s *s3store) delete(uid string) {
	// will write implementation later
	fmt.Println("deleting data from s3")
}

func GetStorage(storeType string) store {
	switch storeType {
	case "file":
		return NewFileStore()
	case "s3":
		return NewS3Store()
	default:
		return NewFileStore()
	}
}

func RunStoreExample() {
	s := GetStorage("file")
	s.write("hello, world", "123")
	s.read("123")
	s.update("this is updated data", "123")
	s.delete("123")

	s = GetStorage("s3")
	s.write("write this to s3", "123")
	s.read("123")
	s.update("updated data", "123")
	s.delete("123")
}
