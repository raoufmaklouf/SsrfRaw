package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sync"

	"github.com/tomnomnom/rawhttp"
)

var wg sync.WaitGroup

func main() {
	path := flag.String("path", "http://omeg0ivn7k95wloyezah7zcqghm7aw.burpcollaborator.net", "oop host")
	regaxSting := flag.String("rg", "k6unx4pudf8k5itoapaxjwzjigz", "regax Sting")
	httpMethod := flag.String("hm", "GET", "http Method")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := scanner.Text()
		if isUrl(line) == true {
			wg.Add(1)

			go func() {
				defer wg.Done()
				body, _ := RawRequest(*httpMethod, line, *path)
				if xMatch(*regaxSting, body) == true {
					fmt.Println(line)

				}

			}()

		}

	}
	wg.Wait()

}

func RawRequest(method string, url string, path string) (body string, scode string) {
	var SCODE string
	var BODY string

	req, err := rawhttp.FromURL(method, url)

	if err == nil {
		req.AutoSetHost()
		req.Path = path
		resp, err2 := rawhttp.Do(req)
		if err2 == nil {
			BODY = string(resp.Body())
			SCODE = string(resp.StatusLine())

		}

	}
	return BODY, SCODE

}

func xMatch(rg string, str string) bool {
	match, _ := regexp.MatchString(rg, str)
	if match == true {
		return true
	} else {
		return false

	}

}

func isUrl(url string) bool {
	s := false
	regex1, _ := regexp.MatchString("http", url)
	regex2, _ := regexp.MatchString("://", url)
	if regex1 == true && regex2 == true {

		s = true
	}
	return s
}
