package api

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
	"v2ex/config"
)

const IMGHTML = "{{img}}"
const SelfLoadTag = "asdffertesdgdsgergdgs"

func RestorePicture(s string, t string, img []string, size ...string) string {
	si := ""
	if len(size) > 0 {
		si = "?w=" + size[0]
	}

	t = strings.ReplaceAll(t, "\"", "")

	_con := config.GetConfig()
	for k, v := range img {
		v = _con.Run.UploadServer + v
		if t != "" {
			s = strings.Replace(s, IMGHTML, fmt.Sprintf(`<img src="%s" alt="%s 第%d张">`, v+si, t, k+1), 1)
		} else {
			s = strings.Replace(s, IMGHTML, fmt.Sprintf(`<img src="%s">`, v+si), 1)
		}
	}
	return s
}

func SeparatePicture(_html string) (html string, imgs []string, err error) {

	_html = fmt.Sprintf("<%s>%s</%s>", SelfLoadTag, _html, SelfLoadTag)
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
	doc.Find("br").ReplaceWithHtml("\r\n")
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

			parse, err := url.Parse(href)
			if err != nil {
				return
			}
			if strings.Contains(parse.Host, _con.Run.SiteHost) {
				selection.ReplaceWithHtml(fmt.Sprintf(`<a href="%s">%s</a>`, parse.String(), selection.Text()))
			} else {
				selection.ReplaceWithHtml(fmt.Sprintf(`<a href="/jump-address?u=%s" rel="nofollow"  target="_blank" >%s</a>`, href, selection.Text()))
			}

		} else {
			selection.ReplaceWithHtml(selection.Text())
		}
	})
	//处理img 标签
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")
		if strings.HasPrefix(src, _con.Run.UploadServer) {

			p_url, err := url.Parse(src)
			if err != nil {
				selection.Remove()
			}
			selection.ReplaceWithHtml(IMGHTML)
			imgs = append(imgs, p_url.Path)
			//imgs = append(imgs, strings.ReplaceAll(src, _con.Run.UploadServer, ""))
		} else {
			selection.Remove()
		}
	})
	if doc.Find("p").Size() <= 2 {
		doc.Find("div").Each(func(i int, selection *goquery.Selection) {
			ret, e := selection.Html()
			if e != nil {
				selection.Remove()
				return
			}
			selection.ReplaceWithHtml("<p>" + ret + "</p>")
		})
	}

	//删除 javascript 标签
	doc.Find("script").Remove()
	//删除为空的p标签
	doc.Find("p").Each(func(i int, selection *goquery.Selection) {

		text := selection.Text()
		if strings.Contains(text, "----占位符----") {
			selection.Remove()
			return
		}

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
