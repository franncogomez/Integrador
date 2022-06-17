package avatar

//Para codificar la información:
type CriptoEncoder interface {
	Encodeinformation(name string) (encodeInformation []byte, err error)
}

//Para generar la imagen:
type ImagenGenerator interface {
	GenerateImg(encodeInformation []byte) error
}

//Estructura para llamar al encoder y la imagen:
type Generator struct {
	encoder CriptoEncoder
	imagen  ImagenGenerator
}

//Función que usa el hasheo de la información y la generación de imagen por defecto,
//es decir la que ya está programada en el encoder y en images:
func Defaulavatargenerator() *Generator {
	return &Generator{
		encoder: &encoder.MD5Encoder{},
		imagen:  imagen.GeneratorOne{},
	}
}

//Para que el usuario suba como quiere encodear la información y generar la imagen:
func DefaulCustom(encoder CriptoEncoder, imagen ImagenGenerator) *Generator {
	return &Generator{
		encoder: CriptoEncoder,
		imagen:  ImagenGenerator,
	}
}

//Service: contiene las funciones para generar el avatar:
type Service struct {
	enconder  CriptoEncoder
	generator ImagenGenerator
}

//Information: contiene la información que se va a encodear:
type Informacion struct {
	name string
}

//Función que genera el Avatar:
func (s *Service) GeneraryGuardaAvatar(name string) error {
	enconder, _ := Encodeinformation()
	generator := GenerateImg()
	return nil
}
