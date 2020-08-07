package manage

import "v2ex/config"

var _con = config.GetConfig()

func Init() {
	_con = config.GetConfig()
}
