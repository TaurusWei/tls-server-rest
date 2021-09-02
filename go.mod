module tls-server-rest

go 1.13

replace github.com/tjfoc/gmsm v1.4.1 => ./internal/github.com/tjfoc/gmsm

replace github.com/tjfoc/gmtls v1.2.1 => ./internal/github.com/tjfoc/gmtls

replace github.com/gin-gonic/gin => ./third_party/github.com/gin-gonic/gin

replace github.com/gin-contrib/sse => ./third_party/github.com/gin-contrib/sse

replace github.com/gin-contrib/pprof v1.3.0 => ./third_party/github.com/gin-contrib/pprof

replace github.com/hyperledger/fabric => ./third_party/github.com/hyperledger/fabric

replace github.com/tauruswei/go-netsign => ./third_party/github.com/tauruswei/go-netsign

replace github.com/unrolled/secure => ./third_party/github.com/unrolled/secure

replace github.com/spf13/viper v1.8.1 => github.com/spf13/viper v1.7.1

require (
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/hyperledger/fabric v1.4.12
	github.com/jmoiron/sqlx v0.0.0-20180124204410-05cef0741ade
	github.com/lib/pq v1.10.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.8.1
	github.com/thedevsaddam/gojsonq v2.3.0+incompatible
	github.com/tjfoc/gmsm v1.4.1
	github.com/tjfoc/gmtls v1.2.1
	github.com/unrolled/secure v1.0.9
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
)
