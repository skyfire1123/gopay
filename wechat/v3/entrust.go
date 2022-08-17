package wechat

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// V3ContractPaymentNotify 预扣费通知API
//	Code = 0 is success
//	https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_10.shtml
func (c *ClientV3) V3ContractPaymentNotify(ctx context.Context, contractID string, bm gopay.BodyMap) (wxRsp *ContractPaymentNotifyRsp, err error) {
	uri := fmt.Sprintf(v3ContractPaymentNotify, contractID)
	authorization, err := c.authorization(MethodPost, uri, bm)
	if err != nil {
		return nil, err
	}
	res, si, _, err := c.doProdPost(ctx, bm, uri, authorization)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusNoContent {
		wxRsp = &ContractPaymentNotifyRsp{Code: Success}
	} else {
		wxRsp = &ContractPaymentNotifyRsp{Code: res.StatusCode}
	}
	return wxRsp, c.verifySyncSign(si)
}
