package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type gsmCode string
type siteId string

type Url struct {
	gc gsmCode
}

type Api struct {
	baseUrl string
	// clientId string
	// password string
	// basic    string
}

type Endpoint struct {
	urlTpl string
	resp   SiteInfoResp
}

type SiteInfoResp struct {
	SiteId    siteId
	UserCount int
}

func (u Url) SiteInfo(si siteId) string {
	return fmt.Sprintf("/%s/%s/info", u.gc, si)
}

func (u Url) Health() string {
	return fmt.Sprintf("/%s/health", u.gc)
}

type Fetcher interface {
	Fetch() ([]byte, error)
}

func (a Api) Fetch(path string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", a.baseUrl+path, nil)

	req.Header.Add("Basic", `olala`)
	res, err := client.Do(req)

	// res, err := http.Get(a.baseUrl + url)

	if err != nil {
		log.Println(err)

		return []byte{}, err
	}

	resp, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	return resp, nil
}

func main() {
	fmt.Println("vim-go")
}
