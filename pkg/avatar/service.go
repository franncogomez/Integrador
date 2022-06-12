package avatar

import (
	"crypto/md5"
	"fmt"
	pkg encoder "pkg\avatar\encoder"
	"Repositorio_Local\pkg\avatar\images"
)

//Para codificar la información:
type criptoEncoder interface {
	EncodeInformation (strInformacion string) (encodeInformation []bytes, err error)
}

//Para generar la imagen:
type imagenGenerator interface {
	GenerateImg (encodeInformation bytes[]) error
}

//Service: contiene las funciones para generar el avatar:
type Service struct {
	enconder criptoEncoder
	generator imagenGenerator
}

//Information: contiene la información que se va a encodear:
type Informacion struct{
	nombre string
}

func (s *Service) GeneraryGuardaAvatar (information Informacion) error{
	s.enconder: encodeInformation[]bytes
	s.generator: 
	return nil
}