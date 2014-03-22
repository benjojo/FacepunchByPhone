package main

import (
	"fmt"
	gq "github.com/PuerkitoBio/goquery"
	"strings"
)

type Thread struct {
	ID         string
	ThreadName string
}

func GetSectionThreads(sectionid int) (Threads []Thread, e error) {
	doc, e := gq.NewDocument(fmt.Sprintf("https://facepunch.com/forumdisplay.php?f=%d", sectionid))
	if e != nil {
		return Threads, fmt.Errorf("Could not load section page")
	}
	Threads = make([]Thread, 0)
	doc.Find(".title").Each(func(i int, s *gq.Selection) {
		if i > 9 {
			return
		}
		t := Thread{}
		t.ThreadName = s.Text()
		URL, exists := s.Attr("href")
		if !exists {
			return
		}

		t.ID = strings.Split(URL, "=")[len(strings.Split(URL, "="))-1]
		Threads = append(Threads, t)
	})

	return Threads, e
}
