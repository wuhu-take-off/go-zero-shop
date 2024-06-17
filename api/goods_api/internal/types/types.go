// Code generated by goctl. DO NOT EDIT.
package types

type GoodsAddReq struct {
	GoodsName     string                  `json:"goodsName"`
	GoodsType     int32                   `json:"goodsType"`
	Status        int32                   `json:"status"`
	Description   *string                 `json:"description,omitempty"`
	Specification []*SpecificationAddList `json:"specification"`
}

type GoodsAddResp struct {
}

type GoodsDelReq struct {
	GoodsId int32 `json:"goodsId"`
}

type GoodsDelResp struct {
}

type GoodsList struct {
	MerchantName string               `json:"merchantName"`
	GoodsId      int32                `json:"goodsId"`
	GoodsName    string               `json:"goodsName"`
	GoodsTypeId  int32                `json:"goodsTypeId"`
	Status       int32                `json:"status"`
	List         []*SpecificationList `json:"list"`
}

type GoodsListReq struct {
	Status      int32   `json:"status,omitempty"`
	GoodsTypeId *int32  `json:"goodsTypeId,omitempty"`
	MerchantId  *int32  `json:"merchantId,omitempty"`
	GoodsName   *string `json:"goodsName,omitempty"`
	GoodsID     *int32  `json:"goodsId,omitempty"`
	Page        int32   `json:"page"`
	Limit       int32   `json:"limit"`
}

type GoodsListResp struct {
	Count  int32        `json:"count"`
	Limit  int32        `json:"limit"`
	Number int32        `json:"number"`
	Page   int32        `json:"page"`
	List   []*GoodsList `json:"list"`
}

type GoodsTypeList struct {
	GoodsTypeId   int32  `json:"goodsTypeId"`
	GoodsTypeName string `json:"goodsTypeName"`
}

type GoodsTypeListReq struct {
}

type GoodsTypeListResp struct {
	Count int32            `json:"count"`
	List  []*GoodsTypeList `json:"list"`
}

type GoodsUpDateReq struct {
	GoodsId     int32                `json:"goodsId"`
	GoodsName   *string              `json:"goodsName,omitempty"`
	GoodsTypeId *int32               `json:"goodsTypeId,omitempty"`
	Status      int32                `json:"status"`
	Description *string              `json:"description,omitempty`
	List        []*SpecificationList `json:"specification,omitempty"`
}

type GoodsUpDateResp struct {
}

type SpecificationAddList struct {
	Size      string `json:"size"`
	Inventory int32  `json:"inventory"`
	Score     int32  `json:"score"`
}

type SpecificationList struct {
	SpecificationId  int32  `json:"specificationId"`
	Size             string `json:"size"`
	DisplayInventory int32  `json:"displayInventory"`
	RealInventory    int32  `json:"realInventory"`
	Score            int32  `json:"score"`
}