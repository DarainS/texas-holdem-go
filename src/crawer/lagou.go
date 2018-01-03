package crawer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
)

const (
	LagouBaseUrl = "https://www.lagou.com"
)

func BuildRequestUrl(searchword, year, city string) string {
	r := "https://www.lagou.com/jobs/list_" + searchword + "?px=default&gj=" + year + "&city=" + city
	return r
}

func BuildAjaxUrl(year, city string) string {
	r := "https://www.lagou.com/jobs/positionAjax.json?gj=" + year + "&px=default&city=" + city + "&needAddtionalResult=false&isSchoolJob=0"
	return r
}

func T1() string {
	url := BuildAjaxUrl("3年及以下", "深圳")
	doc, err := buildRequest("go", 1, url)
	if err != nil {
		log.Fatal(err)
	}
	s := doc.Text()
	//s = strings.Replace(s, " ", "", -1)
	fmt.Println(s)
	return s
}

func BuildRequestTest() {
	BuildRequest("go", "3年及以下", "深圳", 2)
}

func BuildRequest(keyword, year, city string, page int) {
	// Request (POST https://www.lagou.com/jobs/positionAjax.json?gj=3%E5%B9%B4%E5%8F%8A%E4%BB%A5%E4%B8%8B&px=default&city=%E6%B7%B1%E5%9C%B3&needAddtionalResult=false&isSchoolJob=0)

	params := url2.Values{}
	params.Set("isFirst", "true")
	params.Set("pn", "2")
	params.Set("kw", "go")
	body := strings.NewReader(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://www.lagou.com/jobs/positionAjax.json?gj=3%E5%B9%B4%E5%8F%8A%E4%BB%A5%E4%B8%8B&px=default&city=%E6%B7%B1%E5%9C%B3&needAddtionalResult=false&isSchoolJob=0", body)
	//parseFormErr := req.ParseForm()
	//if parseFormErr != nil {
	//	fmt.Println(parseFormErr)
	//}
	// Headers
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Origin", "https://www.lagou.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36")
	req.Header.Add("X-Anit-Forge-Token", "None")
	req.Header.Add("Cookie", "user_trace_token=20170319232847-887270c0f2234d6391cc723c7e92b2ec; LGUID=20170319233006-edd19015-0cb8-11e7-9540-5254005c3644; X_HTTP_TOKEN=5682d07ba63746fa458dcc4c177c3dd2; index_location_city=深圳; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1512223194; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1512223194; _ga=GA1.2.968185037.1489937405; JSESSIONID=CFA87235570A3CD48AADBDDE2E9A97E5; _putrc=A1DCA25783F0B612; login=true; unick=邓收港; LGSID=20171230223312-5d0da67a-ed6e-11e7-b8f0-525400f775ce; PRE_UTM=; PRE_HOST=; PRE_SITE=https://www.lagou.com/jobs/list_%E7%AD%96%E5%88%92?city=%E6%B7%B1%E5%9C%B3&cl=false&fromSearch=true&labelWords=&suginput=; PRE_LAND=https://www.lagou.com/jobs/list_%E7%AD%96%E5%88%92?px=default&gj=3%E5%B9%B4%E5%8F%8A%E4%BB%A5%E4%B8%8B&city=%E6%B7%B1%E5%9C%B3; SEARCH_ID=588e5531fddf46f6b60dc3aa439bf9ba; LGRID=20171230223537-b3d7327f-ed6e-11e7-9fc3-5254005c3644; TG-TRACK-CODE=search_code")
	req.Header.Add("Referer", "https://www.lagou.com/jobs/list_go?px=default&gj=3年及以下&city=深圳")
	req.Header.Add("Host", "www.lagou.com")
	req.Header.Add("DNT", "1")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}

func buildRequest(keyword string, page int, url string) (*goquery.Document, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	form := url2.Values{}
	form.Add("first", "false")
	form.Add("pn", string(page))
	form.Add("kd", keyword)
	req.PostForm = form
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	//req.Header.Add("Referer", LagouBaseUrl)
	req.Header.Add("Cookie", "your cookie") // 也可以通过req.Cookie()的方式来设置cookie
	res, err := client.Do(req)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(res)
	return doc, err
}
