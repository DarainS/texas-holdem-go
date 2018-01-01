package crawer

import (
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	"github.com/henrylee2cn/pholcus/common/goquery"         //DOM解析
	. "github.com/henrylee2cn/pholcus/app/spider"           //必需
	"strings"
	"net/http"
	"log"
)

//修改这个为其他岗位的，可以爬取其他岗位的数据

//const positionURL = "https://www.lagou.com/zhaopin/go/?filterOption=3"
const positionURL = "https://www.lagou.com/jobs/list_python?fromSearch=true&labelWords=relative"

func init() {

	header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")

	header.Add("Accept-Encoding", "gzip, deflate, br")

	header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,fr;q=0.4,tr;q=0.2,zh-TW;q=0.2")

	header.Add("Connection", "keep-alive")

	//header.Add("Cookie", "user_trace_token=20170910220432-f801c133-9630-11e7-8e11-525400f775ce; LGUID=20170910220432-f801c565-9630-11e7-8e11-525400f775ce; index_location_city=%E5%85%A8%E5%9B%BD; JSESSIONID=ABAAABAAADEAAFI27EBBC4DCA6B9DBD97414B0004A32D4F; TG-TRACK-CODE=index_navigation; _gat=1; PRE_UTM=; PRE_HOST=; PRE_SITE=https%3A%2F%2Fwww.lagou.com%2Fzhaopin%2Fgo%2F3%2F%3FfilterOption%3D2; PRE_LAND=https%3A%2F%2Fwww.lagou.com%2Fzhaopin%2Fgo%2F4%2F%3FfilterOption%3D3; SEARCH_ID=418a46d847344429b67029bc1470f19c; _gid=GA1.2.1008155537.1505828050; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1505052272,1505828050; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1505830015; _ga=GA1.2.319466696.1505052272; LGSID=20170919220506-8a4a46e3-9d43-11e7-99b2-525400f775ce; LGRID=20170919220655-cb047879-9d43-11e7-99b2-525400f775ce")

	header.Add("DNT", "1")

	header.Add("Host", "www.lagou.com")

	header.Add("Referer", "https://www.lagou.com/")

	header.Add("Upgrade-Insecure-Requests", "1")

	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3112.113 Safari/535.36\"")

	lagou.Register()

}

var lagou = &Spider{
	Name: "拉勾-岗位",
	Description: "拉勾上的全部岗位【https://www.lagou.com】",
	EnableCookie: true,
	NotDefaultField: true,
	RuleTree: lagouRuleTree,
}

var header = http.Header{}

var lagouRuleTree = &RuleTree{
	Root: func(ctx *Context) {
		ctx.AddQueue(&request.Request{
			Url: positionURL,
			TryTimes: 10,
			Rule: "requestList",
			Header: header,
		})
	},

	Trunk: map[string]*Rule{
		"requestList": {
			ParseFunc: func(ctx *Context) {
				header.Set("Referer", ctx.Request.Url)
				nextSelection := ctx.GetDom().Find("div.pager_container").Find("a").Last()
				url, _ := nextSelection.Attr("href")
				if len(url) != 0 && strings.HasPrefix(url, "http") {
					ctx.AddQueue(&request.Request{

						Url: url,

						TryTimes: 10,

						Rule: "requestList",

						Priority: 1,

						Header: header,
					})
				}
				ctx.Parse("outputResult")
			},
		},

		"outputResult": {
			ItemFields: []string{
				"岗位",
				"薪水",
				"工作地点",
				"公司",
				"工作年限",
				"学历",
				"job网址",
				"企业类型",
				"融资状态",
			},

			ParseFunc: func(ctx *Context) {

				dom := ctx.GetDom()

				/*
				dom.Find("div.list_item_top").Each(func(i int, selection *goquery.Selection) {
					log.Fatal("sa")
					jobName := selection.Find("div.p_top").Find("h3").Text()

					city := selection.Find("div.p_top").Find("em").Text()

					city = strings.Split(city, "·")[0]

					salay := selection.Find("div.p_bot").Find("span.money").Text()
					workstring:=selection.Find("div.p_bot").Find("li_b_l").Text()
					workstring=strings.TrimSpace(workstring)
					workstring=strings.Replace(workstring," ","",-1)
					worktime:=strings.Split(workstring,"/")[0]
					studytime:=strings.Split(workstring,"/")[1]
					joburl:=selection.Find("div.p_top").Find(".position_link#href").Text()
					company := selection.Find("div.company").Find("a").Text()
					ctx.Output(map[int]interface{}{
						0: jobName,
						1: salay,
						2: city,
						3: company,
						4:worktime,
						5:studytime,
						6:joburl,
					})

				})
				//*/
				//*
				dom.Find(".con_list_item").Each(func(i int, selection *goquery.Selection) {
					ptop := selection.Find("div.p_top")
					pbot := selection.Find("div.p_bot")
					jobName, find := selection.Attr("data-positionname")
					if !find {
						log.Fatal(string(i) + "th job name not find: " + ctx.GetUrl())
					}
					companyName, find := selection.Attr("data-company")
					if !find {
						log.Fatal(string(i) + "th company name not find: " + ctx.GetUrl())
					}
					salary, find := selection.Attr("data-salary")
					if !find {
						log.Fatal(string(i) + "th salary not find: " + ctx.GetUrl())
					}
					city := ptop.Find("em").Text()
					city = strings.Split(city, "·")[0]
					jobId, find := selection.Attr("data-positionid")
					if !find {
						log.Fatal(string(i) + "thjob id name not find: " + ctx.GetUrl())
					}
					jobUrl := "https://www.lagou.com/jobs/" + jobId + ".html"
					workstring := pbot.Find(".li_b_l").Text()
					studyTime := strings.TrimSpace(strings.Split(workstring, "/")[1])
					workTime := strings.TrimSpace(strings.Split(workstring, "/")[0])
					company := selection.Find(".company")
					companyStr := company.Find(".company_name").Text()
					companyDomain := strings.Split(companyStr, "/")[0]
					rongzi := strings.Split(companyStr, "/")[1]
					ctx.Output(map[int]interface{}{
						0: jobName,
						1: salary,
						2: city,
						3: companyName,
						4: workTime,
						5: studyTime,
						6: jobUrl,
						7: companyDomain,
						8: rongzi,
					})
				})
				//*/
			},
		},
	},
}
