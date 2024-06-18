// Code generated by goctl. DO NOT EDIT.
package types

type Img struct {
	Data []byte `json:"img"`
}

type LoadImgReq struct {
	GoodsId int32 `json:"goodsId"`
}

type LoadImgResp struct {
	Imgs []*Img `json:"imgs"`
}

type NewTokenReq struct {
}

type NewTokenResp struct {
	Token          string `json:"token"`
	ExpirationTime int64  `json:"expirationTime"`
	RefreshTime    int64  `json:"refreshTime"`
}

type SendAuthReq struct {
	Email string `json:"email"`
}

type SendAuthResp struct {
	OK bool `json:"ok"`
}

type UploadImgReq struct {
	GoodsId int32 `form:"goodsId"`
}

type UploadImgResp struct {
	URL string `json:"url"`
}
