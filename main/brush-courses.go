package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	courseId := "212"
	cookieStr := "cookie"

	fmt.Println("courseId:" + courseId)
	for i := 4712; i <= 4712; i++ {
		chapterId := strconv.Itoa(i)
		learnTime := getCourse(courseId, chapterId, cookieStr)
		postCourse(courseId, chapterId, cookieStr, learnTime)
	}

}

func getCourse(courseId, chapterId, cookieStr string) string {
	urlPath := "https://www.bjjnts.cn/lessonStudy/" + courseId + "/" + chapterId

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlPath, nil)
	header := req.Header
	header.Set("Connection", "keep-alive")
	header.Set("Pragma", "no-cache")
	header.Set("Cache-Control", "no-cache")
	header.Set("Accept", "*/*")
	header.Set("X-Requested-With", "XMLHttpRequest")
	header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	header.Set("Origin", "https://www.bjjnts.cn")
	header.Set("Sec-Fetch-Site", "same-originn")
	header.Set("Sec-Fetch-Mode", "cors")
	header.Set("Sec-Fetch-Dest", "empty")
	header.Set("Referer", urlPath)
	header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	header.Set("Cookie", cookieStr)

	resp, err := client.Do(req)
	if err != nil {
		print(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	reg := regexp.MustCompile(`var duration = parseInt\("[0-9]+"\)`)
	regNum := regexp.MustCompile(`[0-9]+`)

	return regNum.FindString(reg.FindString(string(body)))
}

func postCourse(courseId, chapterId, cookieStr, learnTime string) {
	urlPath := "https://www.bjjnts.cn/addstudentTaskVer2/" + courseId + "/" + chapterId
	urlReferer := "https://www.bjjnts.cn/lessonStudy/" + courseId + "/" + chapterId
	client := &http.Client{}

	values := url.Values{}
	values.Add("learnTime", learnTime)

	req, _ := http.NewRequest("POST", urlPath, strings.NewReader(values.Encode()))
	header := req.Header
	header.Set("Connection", "keep-alive")
	header.Set("Pragma", "no-cache")
	header.Set("Cache-Control", "no-cache")
	header.Set("Accept", "*/*")
	header.Set("X-Requested-With", "XMLHttpRequest")
	header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	header.Set("Origin", "https://www.bjjnts.cn")
	header.Set("Sec-Fetch-Site", "same-origin")
	header.Set("Sec-Fetch-Mode", "cors")
	header.Set("Sec-Fetch-Dest", "empty")
	header.Set("Referer", urlReferer)
	header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	header.Set("Cookie", cookieStr)

	resp, err := client.Do(req)
	if err != nil {
		print(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	m := make(map[string]string)
	_ = json.Unmarshal(body, &m)
	fmt.Println("chapterId:", chapterId, ", learnTime:", learnTime, ", msg:", m["msg"])
}
