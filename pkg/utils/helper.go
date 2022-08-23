package utils

import (
	"io/ioutil"
	"os"
)

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func MustLoadFile(path string) []byte {
	f, err := os.Open(path)
	checkErr(err)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	checkErr(err)
	return b
}

// 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func LoadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}
