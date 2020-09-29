package question

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"time"
	"v2ex/app/api"
	"v2ex/model"
)

type _comment_root_add struct {
	DID  model.DIDTYPE `json:"did" validate:"gt=0" comment:"文章ID"`
	Txt  string        `json:"txt"`
	Text string        `validate:"min=10,max=100000000" comment:"数据"`
}

func commentRootAdd(c *gin.Context) {
	//获取用户信息
	user_info := api.GetNowUserInfo(c)
	if user_info.MID == 0 {
		result_json := c_code.V1GinError(101, "非法操作")
		c.JSON(200, result_json)
		return
	}
	_f := _comment_root_add{}
	c.BindJSON(&_f)
	_f.Text = c_code.RemoveHtmlTag(_f.Txt)
	validator := api.VerifyValidator(_f)
	if validator != "" {
		result_json := c_code.V1GinError(102, validator)
		c.JSON(200, result_json)
		return
	}
	//验证是否回答过此问题
	is_reply := model.CommentQuestionRoot{}
	err := mc.Table(is_reply.Table()).Where(bson.M{"mid": user_info.MID, "did": _f.DID}).FindOne(&is_reply)
	if err != nil {
		result_json := c_code.V1GinError(103, "查询失败")
		c.JSON(200, result_json)
		return
	}

	if is_reply.ID.Hex() != mc.Empty {
		result_json := c_code.V1GinError(104, "不允许重复回答")
		c.JSON(200, result_json)
		return
	}

	//分离数据
	_html, _imgs, err2 := api.SeparatePicture(_f.Txt)
	if err2 != nil {
		result_json := c_code.V1GinError(105, "html解析错误")
		c.JSON(200, result_json)
		return
	}

	//检测did 是否存在
	index := model.DataIndex{}
	mc.Table(index.Table()).Where(bson.M{"did": _f.DID, "d_type": model.DTYPEQuestion}).FindOne(&index)
	if index.DID == 0 {
		result_json := c_code.V1GinError(106, "未找到该问题")
		c.JSON(200, result_json)
		return
	}
	//通过 插入数据库
	comment_root := model.CommentQuestionRoot{
		ID:     primitive.NewObjectID(),
		MID:    user_info.MID,
		RC:     0,
		DID:    index.DID,
		ZanLen: 0,
	}
	err = mc.Table(comment_root.Table()).Insert(comment_root)
	if err != nil {
		result_json := c_code.V1GinError(107, "写入失败")
		c.JSON(200, result_json)
		return
	}

	//写入数据存储表
	comment_text := model.CommentQuestionText{
		ID:          comment_root.ID,
		Text:        _html,
		Zan:         nil,
		Img:         _imgs,
		ReleaseTime: time.Now(),
	}
	//写进数据表中
	err = mc.Table(comment_text.Table()).Insert(comment_text)

	if err != nil {
		result_json := c_code.V1GinError(108, "写入失败")
		mc.Table(comment_root.Table()).Where(bson.M{"_id": comment_root.ID}).DelOne()
		c.JSON(200, result_json)
		return
	}
	ref := c.GetHeader("Referer")
	_u := regexp.MustCompile(`/`+model.UrlTagQuestionReply+`/[\w|\s]{24}`).ReplaceAllString(ref, "")
	_u += "/" + model.UrlTagQuestionReply + "/" + comment_root.ID.Hex()
	result_json := c_code.V1GinSuccess(comment_root.ID, "添加成功", _u)
	//评论字段加 1
	mc.Table(index.Table()).Where(bson.M{"_id": index.ID}).UpdateOne(bson.M{"rc": index.RC + 1, "ct": time.Now().Unix()})

	model.Movement(user_info.MID, index.MID).AddQuestionCommentRoot(comment_root)

	c.JSON(200, result_json)
	return
}
