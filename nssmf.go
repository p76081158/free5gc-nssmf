package nssmf

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/p76081158/free5gc-nssmf/yamlparse"
	//"github.com/p76081158/free5gc-nssmf/bindata"
)

func GetgNBinfo() (gnb_ip string, gnb_n3_ip_B string) {

	return gnb_ip, gnb_n3_ip_B
}

func ApplyNetworkSlice(snssai string, gnb_ip string, gnb_n3_ip_B string, ngci string, cpu int) {
	arg := snssai + " " + gnb_ip + " " + gnb_n3_ip_B + " " + ngci + " " + strconv.Itoa(cpu)
	// test, err := bindata.Asset("shell-script/slice-create.sh")
	// if err != nil {
	// 	// Asset was not found.
	// }
	slice_cmd := "shell-script/slice-create.sh " + arg
	input_cmd := slice_cmd
	cmd := exec.Command("/bin/sh", "-c", input_cmd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}

func DeleteNetworkSlice(snssai string) {
	arg := snssai
	slice_cmd := "shell-script/slice-delete.sh " + arg
	input_cmd := slice_cmd
	cmd := exec.Command("/bin/sh", "-c", input_cmd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}

func AdjustNetworkSlice(snssai string, NF string, cpu int) {
	yamlparse.ModifyCPU(snssai, NF, cpu)
}

// func main() {
// 	//ApplyNetworkSlice("0x01010203", "192.168.72.50", "200", "466-01-000000010")
// 	DeleteNetworkSlice("0x01010203")
// }
// 0x01010203
// 192.168.72.50
// 10.200.100.3 (200)
// 466-01-000000010

// apply_network_slice("0x01010203", "192.168.72.50", "200", "466-01-000000010")
// delete_network_slice("0x01010203"