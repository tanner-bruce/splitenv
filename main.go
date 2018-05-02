package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile(".env")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(strings.NewReader(string(dat)))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		t := sc.Text()
		if len(t) == 0 {
			continue
		}
		re := regexp.MustCompile("(.*?)=(.*)")
		ms := re.FindAllSubmatch([]byte(t), -1)

		if len(ms) == 0 {
			fmt.Println("skipping string ", string(t))
			continue
		}

		if len(ms[0]) < 2 {
			fmt.Println("skipping string ", string(t))
			continue
		}

		k := ms[0][1]
		v := ms[0][2]

		err := ioutil.WriteFile(string(k), v, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}
