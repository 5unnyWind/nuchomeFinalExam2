package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerKitoBio/goquery"
)

//GetHTMLBody can get HTML.
func GetHTMLBody(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
	req.Header.Add("Referer", url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	html := body
	return html
}

func main() {
	f, err := os.Create("./blog.txt")
	if err != nil {
		log.Fatalln(err)
	}
	for i := 1; i <= 7; i++ {
		body := GetHTMLBody("https://blog.lenconda.top/page/" + strconv.Itoa(i))
		html := bytes.NewReader(body)
		dom, err := goquery.NewDocumentFromReader(html)
		if err != nil {
			log.Fatalln(err)
		}
		dom.Find("#index > main > article.post > div > div").Each(func(j int, selection *goquery.Selection) {
			f.WriteString(selection.Text()+"\n\n\n")
		})
	}
	log.Print("已写入blog.txt")

}
