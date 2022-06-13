package avatar

import (
	"crypto/md5"
	"fmt"
	"pkg\avatar\encoder"
	"pkg\avatar\images"
)

//Para codificar la información:
type criptoEncoder interface {
	EncodeInformation (name string) (encodeInformation []bytes, err error)
}

//Para generar la imagen:
type imagenGenerator interface {
	GenerateImg (encodeInformation bytes[]) error
}

//Estructura para llamar al encoder y la imagen:
type Generator struct{
	encoder criptoEncoder
	imagen imagenGenerator
}

//Función que usa el hasheo de la información y la generación de imagen por defecto,
//es decir la que ya está programada en el encoder y en images:
func defaulavatargenerator() *Generator{
	encoder : &Encoder.MD5Encoder{}
	imagen : &Imagen.GeneratoOne{}
}

//Para que el usuario suba como quiere encodear la información y generar la imagen:
func defaulCustom(encoder criptoEncoder, imagen imagenGenerator) *Generator{
	return &Generator{
		encoder: criptoEncoder,
		imagen: imagenGenerator,
	}
}

//Service: contiene las funciones para generar el avatar:
type Service struct {
	enconder criptoEncoder
	generator imagenGenerator
}

//Information: contiene la información que se va a encodear:
type Informacion struct{
	name string
}

//Función que genera el Avatar:
func (s *Service) GeneraryGuardaAvatar (information Informacion) error{
	s.enconder: encodeInformation[]bytes
	s.generator: err
	return nil
}