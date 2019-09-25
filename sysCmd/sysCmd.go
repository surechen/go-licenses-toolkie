package syscmd

import (
	"bufio"
	"fmt"
	"os/exec"
)

//`go list -m all`
func LinuxCmd(exeCmd string) (resultArr []string, err error) {
	cmd := exec.Command("/bin/bash", "-c", exeCmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return resultArr, err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return resultArr, err
	}

	outputBuf := bufio.NewReader(stdout)
	for {
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err)
			} else {
				return resultArr, nil
			}
			//return resultArr, err
		}
		if string(output) == "EOF" {
			continue
		}
		//fmt.Printf("%s\n", string(output))
		resultArr = append(resultArr, string(output))
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return resultArr, err
	}
	return resultArr, err
}
