package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	var isCPUPprof bool //是否开启cpuprofile标志位
	var isMemPprof bool //是否开启内存profile标志位

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		create, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Println("create file failed , error : ", err)
			return
		}
		pprof.StartCPUProfile(create) //往文件中记录cpu profile信息
		defer create.Close()
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 6; i++ {
		go logicCode()
	}

	time.Sleep(20 * time.Second)
	if isMemPprof {
		create, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Println("create file failed , error : ", err)
			return
		}
		pprof.WriteHeapProfile(create)
		defer create.Close()
	}
}

func logicCode() {
	var c chan int //nil
	for {
		select {
		case v := <-c: //阻塞
			fmt.Printf("received from chan , value : %v\n", v)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}
