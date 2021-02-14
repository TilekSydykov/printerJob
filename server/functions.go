package server

import (
	"bufio"
	"os"
	"strings"
)

func parseString(p string) (string, string) {
	p = strings.ReplaceAll(p, "\n", "")
	p = strings.ReplaceAll(p, "\r", "")
	p = strings.ReplaceAll(p, "\f", "")
	slice := strings.Split(p, "=")
	if len(slice) < 2 {
		return "", ""
	}
	return slice[0], slice[1]
}

func RetrieveROM(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}
