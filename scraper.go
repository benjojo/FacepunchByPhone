package main

import (
	"bytes"
	"fmt"
	gq "github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type Thread struct {
	ID         int
	ThreadName string
}

type Post struct {
	Content string
}

func GetThreadPosts(threaddid int) (Posts []Post, e error) {
	doc, e := gq.NewDocument(fmt.Sprintf("https://facepunch.com/showthread.php?t=%d", threaddid))
	if e != nil {
		return Posts, fmt.Errorf("Could not load section page")
	}

	Posts = make([]Post, 0)
	doc.Find(".postcontainer").Each(func(i int, s *gq.Selection) {
		NewPost := Post{}
		NewPost.Content = s.Find(".restore").Text()
		Posts = append(Posts, NewPost)
	})
	return Posts, e
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
		t.ThreadName = fromWindows1252(s.Text())
		URL, exists := s.Attr("href")
		if !exists {
			return
		}
		tid, _ := strconv.ParseInt(strings.Split(URL, "=")[len(strings.Split(URL, "="))-1], 10, 64)
		t.ID = int(tid)

		Threads = append(Threads, t)
	})

	return Threads, e
}

func fromWindows1252(str string) string {
	var arr = []byte(str)
	var buf = bytes.NewBuffer(make([]byte, 512))
	var r rune

	for _, b := range arr {
		switch b {
		case 0x80:
			r = 0x20AC
		case 0x82:
			r = 0x201A
		case 0x83:
			r = 0x0192
		case 0x84:
			r = 0x201E
		case 0x85:
			r = 0x2026
		case 0x86:
			r = 0x2020
		case 0x87:
			r = 0x2021
		case 0x88:
			r = 0x02C6
		case 0x89:
			r = 0x2030
		case 0x8A:
			r = 0x0160
		case 0x8B:
			r = 0x2039
		case 0x8C:
			r = 0x0152
		case 0x8E:
			r = 0x017D
		case 0x91:
			r = 0x2018
		case 0x92:
			r = 0x2019
		case 0x93:
			r = 0x201C
		case 0x94:
			r = 0x201D
		case 0x95:
			r = 0x2022
		case 0x96:
			r = 0x2013
		case 0x97:
			r = 0x2014
		case 0x98:
			r = 0x02DC
		case 0x99:
			r = 0x2122
		case 0x9A:
			r = 0x0161
		case 0x9B:
			r = 0x203A
		case 0x9C:
			r = 0x0153
		case 0x9E:
			r = 0x017E
		case 0x9F:
			r = 0x0178
		default:
			r = rune(b)
		}

		buf.WriteRune(r)
	}

	return strings.Replace(string(buf.Bytes()), "\x00", "", -1)
}
