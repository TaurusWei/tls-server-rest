package model

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/7/15 下午7:42
 */
//Network model
type Network struct {
	Id             int    `db:"id" json:"id"`                          //id
	TenantId       int64  `db:"tenant_id" json:"tenantId"`             //租户id
	Name           string `db:"name" json:"name"`                      //区块链网络名称
	Domain         string `db:"domain" json:"domain"`                  //域
	Desc           string `db:"desc_" json:"desc"`                     //描述
	Type           uint8  `db:"type" json:"type"`                      //网络类型 0-联盟链 1-私有链
	Creator        int64  `db:"creator" json:"creator"`                //创建人员编号
	CreatorName    string `db:"creator_name" json:"creatorName"`       //创建人姓名
	CreateTime     int64  `db:"create_time" json:"createTime"`         //创建时间
	UpdateTime     int64  `db:"update_time" json:"updateTime"`         //修改时间
	MultiChannel   byte   `db:"multi_channel" json:"multiChannel"`     //支持多链 0-否 1-是
	IsInitedLocal  byte   `db:"is_inited_local" json:"isInitedLocal"`  //支持多链 0-否 1-是
	BaseOrgMspID   string `db:"baseorg_mspid" json:"baseOrgMspID"`     //私链模式下，创建网络基础组织 ID
	Status         uint8  `db:"status" json:"status"`                  //状态 0-未初始化 1-正常 2-申请加入中
	CryptoProvider string `db:"crypto_provider" json:"cryptoProvider"` //加密模式  CNCC_GM, GM, SW
}
