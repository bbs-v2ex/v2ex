package api

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
	"v2ex/config"
)

const IMGHTML = "{{--img--}}"
const SelfLoadTag = "asdffertesdgdsgergdgs"

func SeparatePicture(_html string) (html string, imgs []string, err error) {
	_html = fmt.Sprintf("<%s>%s</%s>", SelfLoadTag, _html, SelfLoadTag)
	//_html = strings.ReplaceAll(_html,"<br>","</br>")
	_con := config.GetConfig()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(_html))
	if err != nil {
		return
	}
	defer doc.Clone()

	//处理 h1 标签
	doc.Find("h1").Each(func(i int, selection *goquery.Selection) {
		selection.ReplaceWithHtml(selection.Text())
	})
	doc.Find("br").Remove()
	//处理 a 标签
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")

		if strings.HasPrefix(href, "/jump-address") {
			parse, err := url.Parse("http://127.0.0.1" + href)
			if err != nil {
				selection.ReplaceWithHtml(selection.Text())
				return
			}
			u := parse.Query().Get("u")
			href = u
		}

		if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
			selection.ReplaceWithHtml(fmt.Sprintf(`<a href="/jump-address?u=%s" rel="nofollow"  target="_blank" >%s</a>`, href, selection.Text()))
		} else {
			selection.ReplaceWithHtml(selection.Text())
		}
	})
	//处理img 标签
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")
		if strings.HasPrefix(src, _con.Run.UploadServer) {
			selection.ReplaceWithHtml(IMGHTML)
			imgs = append(imgs, strings.ReplaceAll(src, _con.Run.UploadServer, ""))
		} else {
			selection.Remove()
		}
	})
	//删除 javascript 标签
	doc.Find("script").Remove()
	//删除为空的p标签
	doc.Find("p").Each(func(i int, selection *goquery.Selection) {
		if strings.TrimSpace(selection.Text()) == "" {
			selection.Remove()
		}
		ret, e := selection.Html()
		if e != nil {
			selection.Remove()
			return
		}
		selection.ReplaceWithHtml("<p>" + ret + "</p>")
	})

	html, err = doc.Find(SelfLoadTag).Html()

	return
}