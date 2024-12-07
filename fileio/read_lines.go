package fileio

import (
	"bufio"
	"log"
	"os"
)

func ParseAllLines(path string, parseFn func(string)) (err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseFn(scanner.Text())
	}
	return nil
}
