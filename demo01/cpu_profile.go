package main

import (
	"errors"
	"fmt"
	"gongguowei.com/performace-example/common"
	"gongguowei.com/performace-example/common/op"
	"os"
	"runtime/pprof"
)

var (
	profileName = "cpuProfile.out"
)

func main() {
	f, err := common.CreateFile("", profileName)
	if err != nil {
		fmt.Printf("CPU profile creation error: %v\n\n", err)
		return
	}
	defer f.Close()
	if err := startCPUProfile(f); err != nil {
		fmt.Printf("CPU profile start error: %v\n", err)
		return
	}
	if err = common.Execute(op.CPUProfile, 10); err != nil {
		fmt.Printf("Execute error: %v\n", err)
	}
	stopCPUProfile()
}

func startCPUProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.StartCPUProfile(f)
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}
