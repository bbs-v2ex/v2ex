package common

import "v2ex/config"

func Avatar(u string) string {
	_con := config.GetConfig()
	return _con.Run.UploadServer + u
}
