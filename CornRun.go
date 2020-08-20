package v2ex

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"os"
	"time"
	"v2ex/config"
	"v2ex/model"
	"v2ex/until"
)

// 秒 分 小时 月份中的日期 月份 星期中的日期 年份

//Cron表达式范例：
//
//每隔5秒执行一次：*/5 * * * * ?
//
//每隔1分钟执行一次：0 */1 * * * ?
//
//每天23点执行一次：0 0 23 * * ?
//
//每天凌晨1点执行一次：0 0 1 * * ?
//
//每月1号凌晨1点执行一次：0 0 1 1 * ?
//
//每月最后一天23点执行一次：0 0 23 L * ?
//
//每周星期天凌晨1点实行一次：0 0 1 ? * L
//
//在26分、29分、33分执行一次：0 26,29,33 * * * ?
//
//每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?

func CornRun() {

	_config := config.GetConfig()
	loca_url := fmt.Sprintf("http://127.0.0.1:%d", _config.Run.Port)
	c := cron.New(cron.WithSeconds())
	// 通过AddFunc注册
	c.AddFunc("0 0 */1 * * *", func() {
		//更新网站地图
		c_code.CGet(loca_url + "/update_site_map")
		//删除缓存的图片
		del_tmp_files()
	})

	//清楚会员登录的token
	del_token := false
	//c.AddFunc("0 */10 * * * ?", func() {
	c.AddFunc("* */1  * * * ?", func() {
		if del_token {
			return
		}
		del_token = true
		mc.Table(model.MemberToken{}.Table()).Where(bson.M{"expire": bson.M{"$lt": until.MemberTokenAdd()}}).DelAll()
		del_token = false
	})
	c.Start()
}

func del_tmp_files() {
	_tmp_dir := "/www/wwwroot/www.studyseo.net/view_static/tmp/"
	dir, err := ioutil.ReadDir(_tmp_dir)
	if err != nil {
		return
	}
	for _, f := range dir {
		f.ModTime().Unix()
		fmt.Println(f.ModTime().Add(1 * time.Hour).Format("2006-01-02 15:04:05"))
		if f.ModTime().Add(1*time.Hour).Unix() < time.Now().Unix() {
			os.Remove(_tmp_dir + f.Name())
		}
	}

}
