package avatar

import "crypto/md5"

type MD5Encoder struct{}

func (e *MD5Encoder) EncodeInformation (strInformacion string) (encodeInformation []bytes, err error){
	hash := md5.Sum([]bytes(strInformacion))
	return hash[:], err nil
}
