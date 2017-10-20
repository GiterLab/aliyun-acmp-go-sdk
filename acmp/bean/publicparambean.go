package bean

import "errors"

type formatType string

const (
	XML formatType = "XML"
	JSON                = "JSON"
)

type reginId string

const (
	HANGZHOU reginId ="cn-hangzhou"
)

type PublicParam struct {
	Format           *formatType `json:"format"`
	RegionId         *reginId `json:"region_id"`
	Version          *string `json:"version"`
	AccessKeyId      *string `json:"access_key_id"`
	Signature        *string `json:"signature"`
	SignatureMethod  *string `json:"signature_method"`
	Timestamp        *string `json:"timestamp"`
	SignatureVersion *string `json:"signature_version"`
	SignatureNonce   *string `json:"signature_nonce"`
}

func (this *PublicParam) ToString() (paramstrp *string, err error) {
	if this == nil {
		return nil, errors.New("PublicParam pointer shouldn't be nil")
	}
	if this.RegionId == nil || this.Version == nil || this.AccessKeyId == nil || this.Signature == nil || this.SignatureMethod == nil || this.Timestamp == nil || this.SignatureVersion == nil || this.SignatureNonce == nil {
		return nil, errors.New("PublicParam some perpoties shouldn't be nil")
	}
	var headstr string
	if this.Format != nil {
		headstr = "Format=" + *this.Format + "&"
	} else {
		headstr = ""
	}
	headstr += "RegionId=" + *this.RegionId + "&Version=" + *this.Version + "&AccessKeyId=" + *this.AccessKeyId + "&Signature=" + *this.Signature + "&SignatureMethod=" + *this.SignatureMethod + "&Timestamp=" + *this.Timestamp + "&SignatureVersion=" + *this.SignatureVersion + "&SignatureNonce=" + *this.SignatureNonce
	return &headstr, nil
}
