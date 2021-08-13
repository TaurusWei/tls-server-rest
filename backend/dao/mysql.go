package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
	"tls-server-rest/common/config"
	logger "tls-server-rest/common/log"
)

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/7/15 下午6:55
 */
var sqlDB *sqlx.DB

func OpenSqlDb() {
	var conn string
	// todo DB_MYSQL_CONNECTION
	//conn = os.Getenv("DB_MYSQL_CONNECTION")

	conn = config.GetMysqlConnection()
	for {
		var err error
		sqlDB, err = sqlx.Open("mysql", conn)
		if err != nil {
			logger.Errorf("数据库连接失败：%s", err.Error())
		}
		err = sqlDB.Ping()
		if err != nil {
			logger.Errorf("Open mysql error: %s, wait for 5 seconds to reconnecting…", err.Error())
		} else {
			logger.Info("Open mysql successful…")
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func CloseSqlDb() {
	sqlDB.Close()
}
