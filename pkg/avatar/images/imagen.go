package avatar

import (
	"os"
	"strings"
    "crypto/md5"
	identigo "github.com/Hackdoor-io/identigo"
)

type GeneratorOne struct{}

func (a *GeneratorOne) GenerateImg (encodeInformation []byte) error{
	str,_ := EncodeInformation()
    img := identigo.GenerateFromString(str, 256, 256)

    if err ==nil{
         file, _ := os.Create(str + ".png")
        png.Encode(file, img)
        return err nil
    }else{
        return panic{"No se ha podido generar la imágen"}}
}
