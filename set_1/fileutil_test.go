package set1_test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func scanner(filename string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file), file
}

func getData(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if len(data) == 0 {
		panic(fmt.Sprintf("read file: %s and obtained 0 bytes", fileName))
	}
	return data
}

func getDataTrimNewLine(fileName string) []byte {
	dst := []byte{}
	sc, file := scanner(fileName)
	defer file.Close()
	for sc.Scan() {
		txt := sc.Text()
		dst = append(dst, []byte(txt)...)
	}
	return dst
}

// Reads stringified byte copy pasted into a txt file in the for "xx xx xx"
// where xx is the ineger byte value
func readBytes(fileName string) []byte {
	dst := []byte{}
	scanner, file := scanner(fileName)
	defer file.Close()
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		dst = append(dst, byte(num))
	}
	return dst
}
