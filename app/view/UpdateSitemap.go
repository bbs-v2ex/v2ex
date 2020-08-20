package view

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
	"v2ex/config"
	"v2ex/model"
	"v2ex/until"
)

func UpdateSiteMap(c *gin.Context) {
	update := model.SiteConfig{}.GetUpdateSiteMap()
	_con := config.GetConfig()
	site := _con.Run.SiteMapUrlPreFix
	sitemap_dir := "./__sitemap/"
	isdir, _ := c_code.IsDir(sitemap_dir)
	if !isdir {
		update = true
		os.Mkdir(sitemap_dir, 0666)
	}
	if !update {
		c.String(200, "不更新")
		return
	}
	//os.RemoveAll(sitemap_dir)

	//进行更新sitemap 地图
	//先查询所有文章页
	client := model.DataIndex{}
	page_size := int64(3000)
	order := bson.M{"_id": -1}

	//文章
	pid := primitive.ObjectID{}
	f_index := 1
	link_list := []map[string]string{}

	for {
		a := []model.DataIndex{}
		where := bson.M{
			"d_type": model.DTYPEArticle,
		}
		if pid.Hex() != mc.Empty {
			where["_id"] = bson.M{"$lt": pid}
		}
		mc.Table(client.Table()).Where(where).Projection(bson.M{"_id": 1, "did": 1, "d_type": 1, "t": 1}).Limit(page_size).Order(order).Find(&a)

		if len(a) == 0 {
			break
		}
		for _, v := range a {
			mc.Table(v.InfoArticle.Table()).Where(bson.M{"_id": v.ID}).FindOne(&v.InfoArticle)
			link_list = append(link_list, map[string]string{
				"t":          v.T,
				"loc":        site + model.UrlArticle(v),
				"priority":   "0.3",
				"lastmod":    v.InfoArticle.ReleaseTime.In(until.CST).Format("2006-01-02"),
				"changefreq": "weekly",
			})
			pid = v.ID
		}
		f_name := fmt.Sprintf("/a_%d", f_index)

		f_index++

		//xml
		content := RenderGetContent("sitemap/xml.html", gin.H{"list": link_list})
		content = regexp.MustCompile(`^&lt;\?`).ReplaceAllString(content, "<?")
		ioutil.WriteFile(sitemap_dir+"/"+f_name+".xml", []byte(content), 0666)

		//txt
		content = RenderGetContent("sitemap/text.html", gin.H{"list": link_list})
		content = regexp.MustCompile(`^[\r\n|\n]+`).ReplaceAllString(content, "")
		ioutil.WriteFile(sitemap_dir+"/"+f_name+".txt", []byte(content), 0666)

		//html
		content = RenderGetContent("sitemap/html.html", gin.H{"list": link_list})
		ioutil.WriteFile(sitemap_dir+"/"+f_name+".html", []byte(content), 0666)
		link_list = []map[string]string{}
	}

	//提问
	pid = primitive.ObjectID{}
	f_index = 1
	link_list = []map[string]string{}

	for {
		a := []model.DataIndex{}
		where := bson.M{
			"d_type": model.DTYPEQuestion,
		}
		if pid.Hex() != mc.Empty {
			where["_id"] = bson.M{"$lt": pid}
		}
		mc.Table(client.Table()).Where(where).Projection(bson.M{"_id": 1, "did": 1, "d_type": 1, "t": 1}).Limit(page_size).Order(order).Find(&a)

		if len(a) == 0 {
			break
		}

		for _, v := range a {
			mc.Table(v.InfoQuestion.Table()).Where(bson.M{"_id": v.ID}).FindOne(&v.InfoQuestion)
			link_list = append(link_list, map[string]string{
				"t":          v.T,
				"loc":        site + model.UrlQuestion(v),
				"priority":   "0.3",
				"lastmod":    v.InfoQuestion.ReleaseTime.In(until.CST).Format("2006-01-02"),
				"changefreq": "weekly",
			})
			pid = v.ID
		}
		f_name := fmt.Sprintf("/q_%d", f_index)
		f_index++

		//xml
		content := RenderGetContent("sitemap/xml.html", gin.H{"list": link_list})
		content = regexp.MustCompile(`^&lt;\?`).ReplaceAllString(content, "<?")
		ioutil.WriteFile(sitemap_dir+"/"+f_name+".xml", []byte(content), 0666)

		//txt
		content = RenderGetContent("sitemap/text.html", gin.H{"list": link_list})
		content = regexp.MustCompile(`^[\r\n|\n]+`).ReplaceAllString(content, "")
		ioutil.WriteFile(sitemap_dir+"/"+f_name+".txt", []byte(content), 0666)

		//html
		content = RenderGetContent("sitemap/html.html", gin.H{"list": link_list})
		ioutil.WriteFile(sitemap_dir+"/"+f_name+".html", []byte(content), 0666)
		link_list = []map[string]string{}
	}

	//会员
	pid = primitive.ObjectID{}
	f_index = 1
	link_list = []map[string]string{}

	for {
		a := []model.Member{}
		where := bson.M{}
		if pid.Hex() != mc.Empty {
			where = bson.M{"_id": bson.M{"$lt": pid}}
		}
		mc.Table(model.Member{}.Table()).Where(where).Limit(page_size).Order(order).Find(&a)

		if len(a) == 0 {
			break
		}
		for _, v := range a {

			link_list = append(link_list, map[string]string{
				"t":          v.UserName,
				"loc":        site + model.UrlMember(v),
				"priority":   "0.3",
				"lastmod":    time.Now().In(until.CST).Format("2006-01-02"),
				"changefreq": "weekly",
			})
			pid = v.ID
		}
		f_name := fmt.Sprintf("/member_%d", f_index)
		f_index++

		//xml
		content := RenderGetContent("sitemap/xml.html", gin.H{"list": link_list})
		content = regexp.MustCompile(`^&lt;\?`).ReplaceAllString(content, "<?")
		ioutil.WriteFile(sitemap_dir+"/"+f_name+".xml", []byte(content), 0666)

		//txt
		content = RenderGetContent("sitemap/text.html", gin.H{"list": link_list})
		content = regexp.MustCompile(`^[\r\n|\n]+`).ReplaceAllString(content, "")
		ioutil.WriteFile(sitemap_dir+"/"+f_name+".txt", []byte(content), 0666)

		//html
		content = RenderGetContent("sitemap/html.html", gin.H{"list": link_list})
		ioutil.WriteFile(sitemap_dir+"/"+f_name+".html", []byte(content), 0666)
		link_list = []map[string]string{}
	}
	//创建索引文件
	dir, err := ioutil.ReadDir(sitemap_dir)
	if err != nil {
		return
	}
	index_list := []map[string]string{}
	for _, v := range dir {
		if v.Name() == "index.xml" {
			continue
		}
		if strings.HasSuffix(v.Name(), ".xml") {
			index_list = append(index_list, map[string]string{
				"loc": fmt.Sprintf("%s/site_map_check/%s", site, v.Name()),
			})
		}
	}
	content := RenderGetContent("sitemap/xml_index.html", gin.H{"list": index_list})
	content = regexp.MustCompile(`^&lt;\?`).ReplaceAllString(content, "<?")
	ioutil.WriteFile(sitemap_dir+"/index.xml", []byte(content), 0666)

	content = RenderGetContent("sitemap/text.html", gin.H{"list": index_list})
	ioutil.WriteFile(sitemap_dir+"/index.txt", []byte(content), 0666)
	//更新时间
	model.SiteConfig{}.SetUpdateSiteMap()
	c.String(200, "更新完成")
}
