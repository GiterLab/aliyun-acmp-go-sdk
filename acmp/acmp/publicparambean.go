package acmp

import "errors"

const (
	XML  string = "XML"
	JSON        = "JSON"
)

const (
	HANGZHOU string = "cn-hangzhou"
)

type PublicParam struct {
	Format           string `json:"format"`
	RegionId         string `json:"region_id"`
	Version          string `json:"version"`
	AccessKeyId      string `json:"access_key_id"`
	SignatureMethod  string `json:"signature_method"`
	Timestamp        string `json:"timestamp"`
	SignatureVersion string `json:"signature_version"`
	SignatureNonce   string `json:"signature_nonce"`
}

func (p *PublicParam) ToStringWithoutSignature() (paramstrp string, err error) {
	if p == nil {
		return "", errors.New("PublicParam pointer shouldn't be nil")
	}
	//|| p.Signature == nil
	if p.RegionId == "" || p.Version == "" || p.AccessKeyId == "" || p.SignatureMethod == "" || p.Timestamp == "" || p.SignatureVersion == "" || p.SignatureNonce == "" {
		return "", errors.New("PublicParam some perpoties shouldn't be nil")
	}
	var headstr string
	if p.Format != "" {
		headstr = "Format=" + p.Format + "&"
	} else {
		headstr = ""
	}
	//"&Signature=" + p.Signature +
	headstr += "RegionId=" + p.RegionId + "&Version=" + p.Version + "&AccessKeyId=" + p.AccessKeyId + "&SignatureMethod=" + p.SignatureMethod + "&Timestamp=" + p.Timestamp + "&SignatureVersion=" + p.SignatureVersion + "&SignatureNonce=" + p.SignatureNonce
	return headstr, nil
}
