package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func GetUrlZipRKN() string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://rkn.gov.ru/opendata/7705846236-OperatorsPD/", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	path := doc.Find(".TblList td").Children().Nodes[9].Attr[0].Val

	return "https://rkn.gov.ru" + path
}

func DownloadZipRKN(url string) (err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	f, err := os.Create("/tmp/data.zip")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)

	return
}

func UnzipRKN() error {

	r, err := zip.OpenReader("/tmp/data.zip")
	if err != nil {
		return err
	}
	defer r.Close()

	outFile, err := os.OpenFile("/tmp/data.xml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, r.File[0].Mode())

	if err != nil {
		return err
	}

	rc, err := r.File[0].Open()
	if err != nil {
		return err
	}

	_, err = io.Copy(outFile, rc)
	outFile.Close()
	rc.Close()

	if err != nil {
		return err
	}

	return nil
}

func main() {
	url := GetUrlZipRKN()
	DownloadZipRKN(url)
	UnzipRKN()
	Parsing()
}
