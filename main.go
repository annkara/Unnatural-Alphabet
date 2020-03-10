package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

const (
	japanese = `[\p{Katakana}|\p{Hiragana}|\p{Han}]`
	alphabet = `[a-zａ-ｚ]`
	p        = japanese + alphabet + japanese
)

var re = regexp.MustCompile(p)

func _main(f string) error {

	fd, err := os.Open(f)
	if err != nil {
		return err
	}
	defer fd.Close()

	sc := bufio.NewScanner(fd)
	for sc.Scan() {
		l := sc.Text()
		if re.MatchString(l) {
			fmt.Println(l)
		}
	}

	return nil
}

func main() {

	file := os.Args[1]
	if err := _main(file); err != nil {
		log.Fatal(err)
	}
}
