package main

import "fmt"
import "log"
import "io/ioutil"
import "gopkg.in/yaml.v2"

func ReadConfFile(file_path string) map[interface{}]interface{} {
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Print("Failed to read file ", file_path, "\nCause: ", err)
		return nil
	}

	conf := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal([]byte(data), &conf)
	if err2 != nil {
		log.Print("Parsing failed for ", file_path, "\nCause: ", err2)
		return nil
	} else {
		log.Print("Parsing sucessful for ", file_path)
		return conf
	}
}

func main() {
	conf := ReadConfFile("conf/config.yml")
	fmt.Println("Var:", conf["somevar"])
	fmt.Println("Array:", conf["somearray"])
	fmt.Println("Map:", conf["somemap"])

	broken := ReadConfFile("conf/broken.yml")
	if broken == nil {
		fmt.Println("Conf is nil :(")
	}
	fmt.Println("Broken:", broken)

	inexistent := ReadConfFile("conf/this_doesnt_exists.yml")
	fmt.Println("Inexistent:", inexistent)
}
