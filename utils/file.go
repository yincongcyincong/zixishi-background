package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"errors"
)

func NewFile() *File {
	return &File{}
}

type File struct {
	
}

// get file contents
func (f *File) GetFileContents(filePath string) (content string, err error) {
	defer func(err *error) {
		e := recover()
		if e != nil {
			*err = fmt.Errorf("%s", e)
		}
	}(&err)
	bytes, err := ioutil.ReadFile(filePath)
	content = string(bytes)
	return
}

// file or path is exists
func (f *File) PathIsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// path is empty
func (f *File) PathIsEmpty(path string) bool {
	fs, e := filepath.Glob(filepath.Join(path, "*"))
	if e != nil {
		return false
	}
	if len(fs) > 0 {
		return false
	}
	return true
}

// is write permission
func (f *File) IsWritable(filename string) (error) {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return errors.New("Error: Write permission denied.")
		}else {
			return err
		}
	}
	file.Close()
	return nil
}

// is read permission
func (f *File) IsReadable(filename string) (error)  {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return errors.New("Error: Read permission denied.")
		}else {
			return err
		}
	}
	file.Close()
	return nil
}

// is read and write permission
func (f *File) IsWriterReadable(file string) (error)  {
	err := f.IsWritable(file)
	if err != nil {
		return err
	}
	err = f.IsReadable(file)
	if err != nil {
		return err
	}

	return nil
}


