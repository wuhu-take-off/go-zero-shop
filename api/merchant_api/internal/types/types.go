// Code generated by goctl. DO NOT EDIT.
package types

type MerchantDelReq struct {
	MerchantId int32 `json:"merchantId"`
}

type MerchantDelResp struct {
}

type MerchantList struct {
	MerchantNme    string `json:"merchantNme"`
	MerchantId     int32  `json:"merchantId"`
	MerchantStatus int32  `json:"merchantStatus"`
	Linkname       string `json:"linkname"`
	Role           int32  `json:"role"`
}

type MerchantListReq struct {
	MerchantStatus *int32  `json:"merchantStatus,omitempty"`
	MerchantName   *string `json:"merchantName,omitempty"`
	MerchantId     *int32  `json:"merchantId,omitempty"`
	Page           int32   `json:"page"`
	Limit          int32   `json:"limit"`
}

type MerchantListResp struct {
	Count  int32           `json:"count"`
	Limit  int32           `json:"limit"`
	Number int32           `json:"number"`
	Page   int32           `json:"page"`
	List   []*MerchantList `json:"list"`
}

type MerchantUpDateReq struct {
	MerchantId     int32   `json:"merchantId"`
	MerchantName   *string `json:"merchantName,,omitempty"`
	MerchantStatus *int32  `json:"merchantStatus,omitempty"`
	Linkname       *string `json:"linkname,omitempty"`
	Phone          *string `json:"phone,omitempty"`
}

type MerchantUpDateResp struct {
}
