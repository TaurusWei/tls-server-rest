package util

import (
	"encoding/base64"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
	"tls-server-rest/common/global"
)

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/7/21 下午4:39
 */
/**
  签名
*/
func Sign(data []byte) ([]byte, error) {
	hash, err := global.Verifier.CSP.Hash(data, &bccsp.SHA256Opts{})
	if err != nil {
		return nil, errors.WithMessagef(err, "计算哈希失败, data = %s ", base64.StdEncoding.EncodeToString(data))
	}
	signatrure, err := global.Verifier.Sign(global.Cert.SubjectKeyId, hash)
	if err != nil {
		return nil, err
	}
	return signatrure, nil
}

/**
  验签
*/
func Verify(sig, data, certBytes []byte) (bool, error) {
	hash, err := global.Verifier.CSP.Hash(data, &bccsp.SHA256Opts{})
	if err != nil {
		return false, errors.WithMessagef(err, "计算哈希失败, data = %s ", base64.StdEncoding.EncodeToString(data))
	}
	cert, err := sm2.ReadCertificateFromMem(certBytes)
	if err != nil {
		return false, errors.WithMessagef(err, "解析签名证书失败, cert = %s", string(certBytes))
	}
	result, err := global.Verifier.Verify(cert.SubjectKeyId, sig, hash)
	if err != nil {
		return false, err
	}
	return result, nil
}
