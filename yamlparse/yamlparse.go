package yamlparse

import (
    "fmt"
    "io/ioutil"
    "log"
    //"strings"
    "strconv"

    "gopkg.in/yaml.v2"
)

func ModifyCPU(sliceID string, NF string, cpu int) {
    var resourceYaml Yaml2GoResource
    t := string('"')
	cpuLimit := t + strconv.Itoa(cpu) + "m" + t
	path := "network-slice/" + sliceID + "/" + NF + "-" + sliceID + "/overlays/" + NF + "-" + sliceID + "-cpu.yaml"
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &resourceYaml)
    if err != nil {
        panic(err)
    }
	fmt.Println(resourceYaml.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu)
	resourceYaml.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu = cpuLimit
	resourceYaml.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu   = cpuLimit
    fmt.Println(resourceYaml.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu)
    out, err := yaml.Marshal(&resourceYaml)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))


}
