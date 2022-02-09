package response

import (
	"encoding/hex"
)

type Type51StatusResponse struct {
	Seq uint16
}

func (r *Type51StatusResponse) Marshal() ([]byte, error) {
	response, _ := hex.DecodeString("02CB5100324C90270151902F015490270159902F815C91268161902F810091268005902F80089028800D302F801091278015902F80189026801D902F002091270025903100289028002D9031003090280035903100389129003D3131004091280045913100489027004D9030005090280055913100589128805D913180609128800190318004912980093031800C902980119130801491288019912E801C9128002191320024902900299132002C9129003190330034902A00393132003C912A004190320044902B00499130000190")

	return response, nil
}