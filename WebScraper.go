package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
        "os"
	"strings"
)

type News struct {
TITLE string `parquet:"name=title, type=UTF8, encoding=PLAIN_DICTIONARY"`
PUBDATE string `parquet:"name=pubDate, type=UTF8, encoding=PLAIN_DICTIONARY"`

}

func main() {
	inputDate  := os.Args[1]
        feed := os.Args[2]
        fmt.Println("The argumnet is =", inputDate)
	blogTitles, err := GetLatestBlogTitles(feed,inputDate)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf(blogTitles)
}

func GetLatestBlogTitles(url string, input string) (string,error) {
	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// Save each .post-title as a list
	titles := ""
//	fw, err := local.NewLocalFileWriter("/Users/praveensubramani/workspace/goLang/flat.parquet")
//	pw, err := writer.NewParquetWriter(fw, new(News), 2)
  //      pw.RowGroupSize = 128 * 1024 * 1024 //128M
    //    pw.CompressionType = parquet.CompressionCodec_SNAPPY

	doc.Find(".media-body.top-news-text").Each(func(i int, s *goquery.Selection) {
		pubDate :=s.Find("span.time-dt").Text()
                
		if (strings.Contains(pubDate,input) == true){
		/*news := News{
		TITLE : s.Text(),
		PUBDATE: pubDate,
		}
		pw.Write(news)	*/
			titles += "- " + s.Text() + "\n"
		}
	})
//	fw.Close()
	return titles, nil
}
