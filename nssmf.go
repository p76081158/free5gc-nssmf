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

func DeployNetworkSlice(snssai string, gnb_ip string, gnb_n3_ip_B string, ngci string, cpu int, core_function_cpu int) {
	arg := snssai + " " + gnb_ip + " " + gnb_n3_ip_B + " " + ngci + " " + strconv.Itoa(cpu) + " " +  strconv.Itoa(core_function_cpu)
	slice_cmd := "shell-script/slice-deploy.sh " + arg
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

func ScaleOut(snssai string, ngci string ) {
	arg := snssai + " " + ngci
	slice_cmd := "shell-script/slice-scaleout.sh " + arg
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

func ApplyNetworkSlice(snssai string, ngci string) {
	arg := snssai + " " + ngci
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
	return
}

func ApplyServiceCpuChange(snssai string, ngci string, cpu int) {
	// kubectl -n free5gc set resources deployment service-466-93-000000010-0x01030203 --limits cpu=200m --requests cpu=200m
	arg := ngci + "-" + snssai + " --limits cpu=" + strconv.Itoa(cpu) + "m --requests cpu=" + strconv.Itoa(cpu) + "m"
	slice_cmd := "kubectl -n free5gc set resources deployment service-" + arg
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
