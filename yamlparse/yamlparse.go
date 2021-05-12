package yamlparse

import (
    //"fmt"
    "io/ioutil"
    "log"
    //"strings"
    "strconv"

    "gopkg.in/yaml.v2"
)

func ModifyCPU(sliceID string, NF string, cpu int) {
    var resourceYaml Yaml2GoResource
	cpuLimit := strconv.Itoa(cpu) + "m"
	path := "network-slice/" + sliceID + "/" + NF + "-" + sliceID + "/overlays/" + NF + "-" + sliceID + "-cpu.yaml"
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &resourceYaml)
    if err != nil {
        panic(err)
    }
	
	resourceYaml.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu = cpuLimit
	resourceYaml.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu   = cpuLimit
}
