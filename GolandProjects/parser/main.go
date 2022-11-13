package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type leroyJSON struct {
	Content []struct {
		Link  string `json:"productLink"`
		Name  string `json:"displayedName"`
		Price struct {
			BathPrice int `json:"main_price"`
		} `json:"price"`
	} `json:"content"`
}

var href []string

func visit(vod []string, site *html.Node) {
	if site.Data == "span" {
		fmt.Println()
	}
	if site.Type == html.ElementNode && site.Data == "a" {
		var prev struct {
			link  string
			price string
		}
		for _, v := range site.Attr {
			if v.Val == "ctlg-img vert-hldr" {
				href = append(href, prev.link)
			}
			prev.link = v.Val
		}
	}
	for c := site.FirstChild; c != nil; c = c.NextSibling {
		visit(vod, c)
	}
}

func getLeroyMerlin(res *excelize.File) {
	leroyBath := leroyJSON{}
	client := &http.Client{}
	var data = strings.NewReader(`{"familyIds":["5abcdbf0-7370-11eb-b55b-3b63a6aba6e4_Opus_Family"],"limit":30,"regionId":"506","facets":[{"id":"06575","values":["JACOB DELAFON","ROCA"]},{"id":"currentStoreSellingPrice","min":30000,"max":55000},{"id":"00256","min":170,"max":180},{"id":"00053","min":70,"max":70}],"suggest":true,"offset":0,"searchMethod":"DEFAULT"}`)
	req, err := http.NewRequest("POST", "https://api.leroymerlin.ru/hybrid/v1/getProducts", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "cookie_accepted=true; dtCookie=v_4_srv_2_sn_559FDB5A3C90959B79733177B99038AE_perc_64544_ol_0_mul_2_app-3Ab82b63450c1d92de_1_rcs-3Acss_0; iap.uid=23c66676dad44156871252e0c3c678ce; fromRegion=34; plp:facet:eligibilityByStores=%D0%98%D1%81%D0%BF%D1%8B%D1%82%D0%B0%D1%82%D0%B5%D0%BB%D0%B5%D0%B9; _flowbox=eadd41c0-6d2b-5aac-ca97-1c7ed73ceff1; user-geolocation=0%2C0; disp_delivery_ab=A; disp_react_aa=2; ggr-widget-test=1; sawOPH=true; lastConfirmedRegionID=506; _regionID=506; qrator_ssid=1667865322.060.I1cwdenYCGaBDRIp-fo947ehambmbom9slj9uvhok6he6c665")
	req.Header.Set("DNT", "1")
	req.Header.Set("Origin", "https://spb.leroymerlin.ru")
	req.Header.Set("Referer", "https://spb.leroymerlin.ru/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("x-api-key", "Yeg8l3zQDwpVNBDTP3q6jM4lQVLW5TTv")
	req.Header.Set("x-request-id", "41dc90d483c855d1167496f63791c7c5")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bodyText, &leroyBath)
	res.SetSheetName("Sheet1", "LeroyMerlin")
	res.SetCellValue("LeroyMerlin", "A1", leroyBath.Content[0].Name)
	res.SetCellValue("LeroyMerlin", "B1", leroyBath.Content[0].Price.BathPrice)
	res.SetCellValue("LeroyMerlin", "C1", "https://spb.leroymerlin.ru"+leroyBath.Content[0].Link)
	res.SetColWidth("LeroyMerlin", "A", "A", 40)
	res.SetColWidth("LeroyMerlin", "B", "B", 9)
	res.SetColWidth("LeroyMerlin", "C", "C", 100)
	fmt.Println("Leroy Merlin successfully parsed!")
}

func getVodopad(res *excelize.File) {
	var vod []string
	urlVodopad := "https://vodopad.ru/catalog/75979/f/price-da48827f-e77a-4eaf-ad5c-79fef5601e3c-from-30000-to-55000/dlina_sm-from-170-to-180/proizvoditel-is-roca-or-jacob_delafon/shirina_sm-from-70-to-70/"
	getBath, _ := http.Get(urlVodopad)
	site, err := html.Parse(getBath.Body)
	if err != nil {
		_ = fmt.Errorf("error parsing vodopad html %s", err)
		os.Exit(0)
	}
	visit(vod, site)
	res.NewSheet("Vodopad")
	for i := 0; i < len(href); i++ {
		res.SetCellValue("Vodopad", "C"+strconv.Itoa(i+1), "https://vodopad.ru/catalog"+href[i])
	}
	//res.SetCellValue("Vodopad", "A1", leroyBath.Content[0].Name)
	//res.SetCellValue("Vodopad", "B1", leroyBath.Content[0].Price.BathPrice)
	res.SetColWidth("Vodopad", "A", "A", 40)
	res.SetColWidth("Vodopad", "B", "B", 9)
	res.SetColWidth("Vodopad", "C", "C", 100)
	fmt.Println("Vodopad successfully parsed!")
}

func main() {
	res := excelize.NewFile()
	//getLeroyMerlin(res)
	getVodopad(res)
	res.SaveAs("Baths.xlsx")
}
