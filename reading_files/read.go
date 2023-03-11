package reading_files

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadStats(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	status, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(status.Name())
	fmt.Println(status.ModTime().Format("15:04:03"))

	defer file.Close()
}

func ReadWholeFile(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(contents))
}

func ReadByLine(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByWord(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByBytes(filename string, size uint8) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	buf := make([]byte, size)

	for {
		n, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(buf[:n]))
	}
}
