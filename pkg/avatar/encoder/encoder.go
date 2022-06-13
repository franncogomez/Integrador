package avatar

import "crypto/md5"

type MD5Encoder struct{}

func (e *MD5Encoder) Encodeinformation (name string) (encodeInformation []byte, err error){
	hash := md5.Sum([]byte (name))
	return hash[:], err
}
