package login

import (
	"net/http"

	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/proto"
)

func (this *Login) MsgLogin(w http.ResponseWriter, req *http.Request, data string) {
	msg := this.decodeMsg(data, &proto.MsgLogin{}).(*proto.MsgLogin)
	common.GetLogger().Debugln("account =", msg.GetAccount())
	common.GetLogger().Debugln("password =", msg.GetPassword())
	common.GetLogger().Debugln("mode =", msg.GetMode())

	accountId, err := this.Derived.OnVerifyAccount(msg.GetAccount(), msg.GetPassword(), msg.GetMode(), msg.GetUserdata())
	if err != nil {
		return
	}

	common.GetLogger().Debugln("accountId =", accountId)
}
