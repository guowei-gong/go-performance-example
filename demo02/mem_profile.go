package main

import (
	"errors"
	"fmt"
	"gongguowei.com/performace-example/common"
	"gongguowei.com/performace-example/common/op"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	profileName = "memProfile.out"
	memProfileRate = 8 // 每在堆内存上平均分配 8 字节, 进行一次采样。
)

func main() {
	f, err := common.CreateFile("", profileName)
	if err != nil {
		fmt.Printf("memory profile creation error: %v\n", err)
		return
	}
	defer f.Close()
	startMemProfile()
	if err = common.Execute(op.MemProfile, 10); err != nil {
		fmt.Printf("execute")
		return
	}
	if err := stopMemProfile(f); err != nil {
		fmt.Printf("memory profile stop error: %v\n", err)
		return
	}
}

func startMemProfile() {
	runtime.MemProfileRate = memProfileRate
}

func stopMemProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.WriteHeapProfile(f)
}