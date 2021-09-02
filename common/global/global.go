package global

import (
	"github.com/hyperledger/fabric/bccsp/verifier"
	"github.com/tjfoc/gmsm/sm2"
)

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/7/21 下午4:43
 */
// Settings
var (
	Verifier  *verifier.BccspCryptoVerifier
	Cert      *sm2.Certificate
	CertBytes []byte
)
