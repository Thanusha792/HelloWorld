package configuration

/*
Config.go for establishing the connection to port
*/
import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

/*
 PortCoonect() for returning connection to port
*/
func PortCoonect() Port {
	yfile, err := ioutil.ReadFile("env/configuration.yaml")

	if err != nil {
		log.Fatal(err)
	}
	var port Port
	err1 := yaml.Unmarshal(yfile, &port)
	if err1 != nil {
		log.Fatal(err1)
	}
	return port
}
