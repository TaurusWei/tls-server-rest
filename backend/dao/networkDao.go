package dao

import "tls-server-rest/backend/model"

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/7/15 下午7:42
 */
var networkv2Columns = " id, tenant_id, name, domain, desc_, type,creator, creator_name, create_time, update_time, multi_channel, baseorg_mspid, type,crypto_provider, status "

func GetNetworkIdByName(networkName string) (int, error) {
	// 在区块链平台创建网络名称时检测网络是否存在是通过domain来查询的；
	// 在保存网络时domain和name是一样的 同时进行了保存；
	// 在修改网络名称后只修改了fb_network表中的name字段，domain并没有同步修改，
	// 生成网络时用的name字段，所以此处用name来进行查询
	var sql = "SELECT id FROM network WHERE name = ? "
	var networkId int
	// 如果Get没有查询到记录会返回 "sql: no rows in result set" 的错误
	err := sqlDB.Get(&networkId, sql, networkName)
	return networkId, err
}

func GetNetwork() ([]model.Network, error) {
	var sql = "SELECT " + networkv2Columns + " from network"
	var networks []model.Network
	err := sqlDB.Select(&networks, sql)
	return networks, err
}

func GetNetworkNameById(id int) ([]model.Network, error) {
	var sql = "SELECT " + networkv2Columns + " from network where id = ?"
	var networks []model.Network
	err := sqlDB.Select(&networks, sql, id)
	return networks, err
}
func GetNetworkByName(name string) ([]model.Network, error) {
	var sql = "SELECT " + networkv2Columns + " from network where name = ?"
	var networks []model.Network
	err := sqlDB.Select(&networks, sql, name)
	return networks, err
}
