package avatar

import(
	"os"
    identigo "github.com/Hackdoor-io/identigo"
)

type GeneratoOne struct{}

(a *GeneratoOne) GenerateImg (encodeInformation bytes[]) error{
	str := encodeIndormation bytes[]
    img := identigo.GenerateFromString(str, 256, 256)

    if err ==nil{
         file, _ := os.Create(str + ".png")
        png.Encode(file, img)
        return err nil
    }else{return panic{"No se ha podido generar la imágen"}}
}
