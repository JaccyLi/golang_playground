package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"jack", "jaccy", "leslie"}
	ages := []int8{12, 34, 25, 18}
	times := [3]int64{}
	var friends [3]string
	friends[0] = "Lisa"
	cpAge := ages[:]
	sliceAge1 := ages[:2]
	sliceAge2 := ages[3:]
	sliceAge3 := ages[1:2]
	sliceAge4 := ages[2:3]

	time1 := int64(time.Now().UnixNano())
	times[0] = time1
	time.Sleep(200 * time.Millisecond)
	time2 := int64(time.Now().UnixNano())
	times[2] = time2
	fmt.Printf("B: time1=%v\ntime2=%v\n", time1, time2)

	times[0], times[1] = times[1], times[2]
	//t1 := time.Now()
	//t2 := time.Unix(0, t1.UnixNano())
	fmt.Printf("first people=%s\n%d\n%v\n", names[0], ages, times)
	fmt.Printf("A: time1=%v\ntime2=%v\n", time1, time2)
	fmt.Printf("cpAge=%v\n", cpAge)
	fmt.Printf("sliceAge1=%v\n", sliceAge1)
	fmt.Printf("sliceAge2=%v\n", sliceAge2)
	fmt.Printf("sliceAge3=%v\n", sliceAge3)
	fmt.Printf("sliceAge4=%v\n", sliceAge4)
	//fmt.Println(t1, t1.UnixNano())
	//fmt.Println(t2, t2.UnixNano())
	fmt.Println(friends)
}
