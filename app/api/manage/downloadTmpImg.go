package manage

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
	"v2ex/config"
)

type _download_temp_img struct {
	U string `json:"u"`
}

func DownloadTempImg(c *gin.Context) {
	_con = config.GetConfig()
	_f := _download_temp_img{}
	c.BindJSON(&_f)
	u := _f.U
	tmp_dir := _con.ExecPath + "/app/" + "view/view_static/tmp/"
	err2 := os.MkdirAll(tmp_dir, 0666)
	if err2 != nil {
		log.Println(err2)
	}
	log.Println(tmp_dir)
	_img_name := c_code.Random(8)
	_img_name = c_code.Md532(time.Now().String() + _img_name)

	re_base64 := `data:([^;]*);base64,`
	if regexp.MustCompile(re_base64).MatchString(u) {
		_u := regexp.MustCompile(re_base64).ReplaceAllString(u, "")
		decode, err2 := c_code.Bs64Decode(_u)
		if err2 != nil {
			result_json := c_code.V1GinError(102, "下载文件失败")
			c.JSON(200, result_json)
			return
		}
		err2 = ioutil.WriteFile(tmp_dir+_img_name, decode, 0600)
		if err2 != nil {
			result_json := c_code.V1GinError(103, "下载文件失败")
			c.JSON(200, result_json)
			return
		}

	} else {
		err := c_code.DownloadFile(u, tmp_dir+_img_name)
		if err != nil {
			result_json := c_code.V1GinError(101, "下载文件失败"+err.Error())
			c.JSON(200, result_json)
			return
		}
	}
	result_json := c_code.V1GinSuccess("/static/tmp/" + _img_name)
	c.JSON(200, result_json)
	return
}
