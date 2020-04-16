package shis

import (
	"bufio"
	"errors"
	"os"
)

func HistFileReader(histories *[]string, histFileName *string) (error, int) {

	histFile, err := os.Open(*histFileName)
	if err != nil {
		return errors.New("file open err: " + *histFileName + " is notfound."), 301
	}

	histLine := bufio.NewScanner(histFile)
	for histLine.Scan() {
		*histories = append(*histories, histLine.Text())
	}
	return nil, 0
}
