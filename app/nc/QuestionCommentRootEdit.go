package nc

import (
	"errors"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

func QuestionCommentRootEdit(txt string, did model.DIDTYPE, mid model.MIDTYPE, answer model.CommentQuestionRoot) error {
	_html, _imgs, err2 := api.SeparatePicture(txt)
	if err2 != nil {
		return errors.New("html解析错误")
	}
	update := bson.M{
		"text": _html,
		"img":  _imgs,
	}
	err := mc.Table(answer.Text.Table()).Where(bson.M{"_id": answer.ID}).UpdateOne(update)
	if err != nil {
		return err
	}
	return nil

}

//if err != nil {
//	result_json := c_code.V1GinError(104, "修改失败")
//	c.JSON(200, result_json)
//	return
//}
//result_json := c_code.V1GinSuccess("修改成功", "", fmt.Sprintf("/%s/%d/%s/%s", model.UrlTagQuestion, _f.DID, model.UrlTagQuestionReply, answer.ID.Hex()))
//c.JSON(200, result_json)
