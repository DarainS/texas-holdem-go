package crawler2

import (
	"github.com/bitly/go-simplejson"
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	. "github.com/henrylee2cn/pholcus/app/spider"           //必需
	"github.com/henrylee2cn/pholcus/common/goquery"         //DOM解析
	"github.com/henrylee2cn/pholcus/logs"
	"net/http"
	"strings"
)

//修改这个为其他岗位的，可以爬取其他岗位的数据

const positionURL = "https://www.lagou.com/zhaopin/go/?filterOption=3"

var posrUrl = "https://www.lagou.com/jobs/positionAjax.json?city=%E6%B7%B1%E5%9C%B3&needAddtionalResult=false&isSchoolJob=0"

func init() {
	initHeader()
	lagou.Register()
}

var headers = map[string]string{
	"X-Requested-With":   "XMLHttpRequest",
	"Connection":         "keep-alive",
	"Accept-Encoding":    "gzip, deflate, br",
	"Content-Type":       "application/x-www-form-urlencoded; charset=UTF-8",
	"Origin":             "https://www.lagou.com",
	"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36",
	"X-Anit-Forge-Token": "None",
	"Cookie":             "user_trace_token=20170319232847-887270c0f2234d6391cc723c7e92b2ec; LGUID=20170319233006-edd19015-0cb8-11e7-9540-5254005c3644; X_HTTP_TOKEN=5682d07ba63746fa458dcc4c177c3dd2; index_location_city=深圳; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1512223194; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1512223194; _ga=GA1.2.968185037.1489937405; JSESSIONID=CFA87235570A3CD48AADBDDE2E9A97E5; _putrc=A1DCA25783F0B612; login=true; unick=Darain; TG-TRACK-CODE=index_search; _gat=1; LGSID=20180107094547-7bb89558-f34c-11e7-a01c-5254005c3644; PRE_UTM=; PRE_HOST=; PRE_SITE=https://www.lagou.com/jobs/list_Java?px=default&gx=%E5%85%A8%E8%81%8C&gj=&isSchoolJob=1&city=%E6%B7%B1%E5%9C%B3; PRE_LAND=https://www.lagou.com/jobs/list_go?city=%E6%B7%B1%E5%9C%B3&cl=false&fromSearch=true&labelWords=&suginput=&isSchoolJob=1; SEARCH_ID=b289084fbd0c4fe0b58f3756dae9a1fd; LGRID=20180107095325-8c71fcd2-f34d-11e7-bfdb-525400f775ce",
	"Referer":            "https://www.lagou.com/jobs/list_游戏策划?labelWords=&fromSearch=true&suginput=",
	"Host":               "www.lagou.com",
	"DNT":                "1",
	"Accept-Language":    "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
	"Accept":             "application/json, text/javascript, */*; q=0.01",
	"X-Anit-Forge-Code":  "0",
}

func initHeader() {
	for key, value := range headers {
		header.Add(key, value)
	}
}

var lagou = &Spider{
	Name:            "拉勾-岗位",
	Description:     "拉勾上的全部岗位【https://www.lagou.com】",
	EnableCookie:    true,
	NotDefaultField: true,
	RuleTree:        lagouRuleTree,
	Keyin:           KEYIN,
}

var header = http.Header{}

var searchUrl2 = "https://www.lagou.com/jobs/list_"

type LagouJobInfo struct {
	PositionName  string
	Salary        string
	Education     string
	WorkYear      string
	CompanyName   string
	City          string
	SecondType    string
	IndustryField string
	positionId    int
	JobUrl        string
	FinanceStage  string
	companyId     int64
	CompanyUrl    string
}

var lagouRuleTree = &RuleTree{
	Root: func(context *Context) {
		keyin := context.GetKeyin()
		keyin = strings.TrimSpace(keyin)
		if len(keyin) == 0 {
			keyin = "游戏策划"
		}
		for _, key := range strings.Split(keyin, "/") {
			context.SetKeyin(key)
			//context.AddQueue(&request.Request{
			//	Method:     "GET",
			//	Url:        searchUrl2 + context.GetKeyin(),
			//	TryTimes:   1,
			//	Rule:       "postResultParse",
			//	Header:     header,
			//	Reloadable: true,
			//})
			context.AddQueue(&request.Request{
				Method:       "POST",
				Url:          posrUrl,
				TryTimes:     10,
				EnableCookie: true,
				Rule:         "postResultParse",
				Priority:     1,
				Header:       header,
				PostData:     "isfirst=false&pn=1" + "&kd=" + context.GetKeyin(),
			})

		}
	},
	Trunk: map[string]*Rule{
		"postResultParse": {
			ParseFunc: func(context *Context) {
				dom := context.GetDom()
				test := dom.Text()
				reader := strings.NewReader(test)
				jsonResult, _ := simplejson.NewFromReader(reader)
				success, err := jsonResult.Get("success").Bool()
				if err != nil || !success {
					logs.Log.Error("postResultParse error")
				}
				pageNo, _ := jsonResult.Get("content").Get("pageNo").Int()
				pageSize, _ := jsonResult.Get("content").Get("pageSize").Int()
				if pageSize < 15 {
					context.AddQueue(&request.Request{
						Method:   "POST",
						Url:      posrUrl,
						TryTimes: 10,
						Rule:     "postResultParse",
						Priority: 1,
						Header:   header,
						PostData: "pn=" + string(pageNo) + "&kd=" + context.GetKeyin(),
					})
				}
			},
		},
		"outputResult": {
			ItemFields: []string{
				"岗位",
				"薪水",
				"工作地点",
				"公司",
			},
			ParseFunc: func(context *Context) {
				dom := context.GetDom()
				dom.Find("div.list_item_top").Each(func(i int, selection *goquery.Selection) {
					jobName := selection.Find("div.p_top").Find("h3").Text()
					city := selection.Find("div.p_top").Find("em").Text()
					city = strings.Split(city, "·")[0]
					salay := selection.Find("div.p_bot").Find("span.money").Text()
					company := selection.Find("div.company").Find("a").Text()
					context.Output(map[int]interface{}{
						0: jobName,
						1: salay,
						2: city,
						3: company,
					})
				})

			},
		},
	},
}
