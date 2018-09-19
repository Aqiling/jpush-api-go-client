package push

import (
	"fmt"
	"net/http"
	"github.com/jukylin/jpush-api-go-client/common"
)

type Cid struct{
	Cid  string `json:"cid"`
}

func NewCidObject(cid string) *Cid {
	return &Cid{cid}
}

func (c *Cid) Validate() error {

	return nil
}

func (c *Cid) Value() string {
	return c.Cid
}

type CidResult struct {
	common.ResponseBase

	// 成功时 msg_id 是 string 类型。。。
	// 失败时 msg_id 是 string 类型。。。
	CidList  []string `json:"cidlist"`
}

// 成功： {"sendno":"18", "msg_id":"1828256757"}
// 失败： {"msg_id": 1035959738, "error": {"message": "app_key does not exist", "code": 1008}}
func (cr *CidResult) FromResponse(resp *http.Response) error {
	cr.ResponseBase = common.NewResponseBase(resp)

	if !cr.Ok() {
		return nil
	}
	return common.RespToJson(resp, &cr)
}

func (pr *CidResult) String() string {
	f := "<CidList> CidList: %v, sendno: \"%s\", \"%s\""
	return fmt.Sprintf(f, pr.CidList)
}