package crawl

import (
	// "fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type BRS struct {
	Judul     string
	Tanggal   string
	Abstraksi string
	Link      string
}

func BRSScrape(url string) BRS {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	judul := doc.Find("#content .judulberita").Text()
	tanggal := doc.Find("#content .tanggal-rilis-item-brs label").Text()
	abstraksi := doc.Find("#content .abstrak-item-brs").Text()
	link, _ := doc.Find("#content #download-brs").Attr("href")

	return BRS{Judul: judul, Tanggal: tanggal, Abstraksi: abstraksi, Link: link}
}

func BRSIndex(url string) []BRS {
	var brs []BRS
	u := url + "/pressrelease.html"
	doc, err := goquery.NewDocument(u)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".items2 tbody tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		link, _ := s.Find(".judul a").Attr("href")
		// fmt.Printf("link : %s\n", link)
		b := BRSScrape(link)

		brs = append(brs, b)
	})
	return brs
}
