package fileio

import (
	"bufio"
	"log"
	"os"
)

func ReadAllLines(path string) (lines []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}
