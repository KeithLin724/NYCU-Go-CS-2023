package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

const URL = "https://www.ptt.cc/bbs/joke/M.1481217639.A.4DF.html"

var maxNumber int

var re = regexp.MustCompile(`^(?:→|推)\s+([^\s]+):\s+(.+)\s+(\d{2}/\d{2}\s+\d{2}:\d{2})$`)

func SplitComment(input string) []string {

	// Find matches
	return re.FindStringSubmatch(input)[1:]
}

func main() {
	flag.IntVar(&maxNumber, "max", 10, "Max number of comments to show")
	flag.Parse()

	// fmt.Println(maxNumber)

	comment := []string{}

	c := colly.NewCollector()
	c.OnHTML("div.push", func(e *colly.HTMLElement) {
		pushContent := strings.TrimSpace(e.Text)
		comment = append(comment, pushContent)
	})

	err := c.Visit(URL)
	if err != nil {
		log.Fatal(err)
	}

	for index, value := range comment[:maxNumber] {
		res := SplitComment(value)
		fmt.Printf("%d. 名字：%s，留言: %s，時間： %s\n", index+1, res[0], res[1], res[2])

	}

}
