package nssmf

import (
	"fmt"
	"os"
	"os/exec"
)

type nssmf struct {
}

func (nf *nssmf) apply_network_slice(snssai string, gnb_ip string, gnb_n3_ip_B string, ngci string) {
	arg := snssai + " " + gnb_ip + " " + gnb_n3_ip_B + " " + ngci
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

func (nf *nssmf) delete_network_slice(snssai string) {
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

// 0x01010203
// 192.168.72.50
// 10.200.100.3 (200)
// 466-01-000000010

// apply_network_slice("0x01010203", "192.168.72.50", "200", "466-01-000000010")
// delete_network_slice("0x01010203"