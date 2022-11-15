package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	// v, _ := mem.VirtualMemory()

	// // almost every return value is a struct
	// fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	// fmt.Println(v)
	//////////
	ID := "000001"
	id, _ := strconv.ParseInt(ID, 10, 32)
	fmt.Println(id)

	cpus, err := cpu.Percent(time.Duration(2)*time.Second, true) //true -> all, false -> single
	if err == nil {
		for i, c := range cpus {
			fmt.Printf("cpu%d : %f%%\n", i, c)
		}
	}
}
