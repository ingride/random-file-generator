package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createBytes(buf *[]byte, count int) {
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {
		intByte := byte(random(0, 9))
		*buf = append(*buf, intByte)
	}
}

func main() {
	minusBS := flag.Int("bs", 0, "Block Size")
	minusCOUNT := flag.Int("count", 0, "Counter")
	numberOfFiles := flag.Int("n", 0, "Number of files")
	flag.Parse()

	if *minusBS < 0 || *minusCOUNT < 0  || *numberOfFiles <= 0 {
		fmt.Println("Count or/and Byte Size and Number of files < 0!")
		os.Exit(-1)
	}

	name := "test"
	root := "site"
	os.MkdirAll(root, os.ModePerm)
	rand.Seed(time.Now().Unix())


	for i:=0; i < *numberOfFiles; i++ {

		filename := root + "/" + name + strconv.Itoa(i)
		_, err := os.Stat(filename)
		if err == nil {
			fmt.Printf("File %s already exists.\n", filename)
			os.Exit(1)
		}

		destination, err := os.Create(filename)
		if err != nil {
			fmt.Println("os.Create:", err)
			os.Exit(1)
		}

		buf := make([]byte, *minusBS)
		buf = nil

		for j := 0; j < *minusCOUNT; j++ {
			createBytes(&buf, *minusBS)
			if _, err := destination.Write(buf); err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			buf = nil
		}
	}
}