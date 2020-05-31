package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/smallnest/goreq"
)

var list string

func getinfo() (content string) {
	resp, err := http.Get("http://aho4ahoaho.main.jp/railway-info/index.php")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(resp.Status)
	}

	//body, _ := ioutil.ReadFile("./list.json")
	json, _ := simplejson.NewJson(body)

	var contentbody string = `"content":"現在の鉄道運行情報です。","embeds":[` +
		strings.ReplaceAll(`{"title":"北海道","description":"replaceH","color":"16543485"},`, "replaceH", json.Get("Hokaido").MustString()) +
		strings.ReplaceAll(`{"title":"東北","description":"replaceT","color":"10833609"},`, "replaceT", json.Get("Tohoku").MustString()) +
		strings.ReplaceAll(`{"title":"関東","description":"replaceKa","color":"14971568"},`, "replaceKa", json.Get("Kanto").MustString()) +
		strings.ReplaceAll(`{"title":"中部","description":"replaceChubu","color":"5512381"},`, "replaceChubu", json.Get("Chubu").MustString()) +
		strings.ReplaceAll(`{"title":"近畿","description":"replaceKi","color":"15607243"},`, "replaceKi", json.Get("Kinki").MustString()) +
		strings.ReplaceAll(`{"title":"中国","description":"replaceChugoku","color":"5934338"},`, "replaceChugoku", json.Get("Chugoku").MustString()) +
		strings.ReplaceAll(`{"title":"四国","description":"replaceS","color":"16407296"},`, "replaceS", json.Get("Shikoku").MustString()) +
		strings.ReplaceAll(`{"title":"九州","description":"replaceKy","color":"8499264"}`, "replaceKy", json.Get("Kyushu").MustString()) +
		`]`
	return contentbody
}

func postmessage(content string) {
	flag.Parse()
	args := strings.Replace(strings.Replace(fmt.Sprint(flag.Args()), "]", "", 1), "[", "", 1)
	contentbody := `{"username":"運行情報bot","avatar_url":"http://aho4ahoaho.main.jp/railway-info/icon.png",` + content + "}"
	resp, _, _ := goreq.New().Post(args).ContentType("application/json").SendMapString(contentbody).End()
	println(resp.Status)
	return
}

func timer() {
	for true {
		t := time.Now()
		fmt.Println(fmt.Sprint(t.Hour()) + "時" + fmt.Sprint(t.Minute()) + "分" + fmt.Sprint(t.Second()) + "秒に" + "ループしました。")
		if t.Hour() == 16 && t.Minute() == 00 {
			for i := 0; i < 3; i++ {
				t := time.Now()
				postmessage(getinfo())
				fmt.Println(fmt.Sprint(t.Hour()) + "時" + fmt.Sprint(t.Minute()) + "分" + fmt.Sprint(t.Second()) + "秒に実行されました。")
				time.Sleep(1 * time.Hour)
			}
			time.Sleep(12*time.Hour + 25*time.Minute)
		}
		if t.Hour() == 7 && t.Minute() == 30 {
			for i := 0; i < 5; i++ {
				t := time.Now()
				postmessage(getinfo())
				fmt.Println(fmt.Sprint(t.Hour()) + "時" + fmt.Sprint(t.Minute()) + "分" + fmt.Sprint(t.Second()) + "秒に実行されました。")
				time.Sleep(30 * time.Minute)
			}
			time.Sleep(6*time.Hour + 25*time.Minute)
		}
		time.Sleep(20 * time.Second)
	}
}

func main() {
	fmt.Println("bot起動しました(^o^)")
	timer()
	fmt.Println("bot終了します。無限ループ文があるのに何故呼ばれたんでしょう？")
}
