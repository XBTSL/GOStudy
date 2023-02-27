package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

//go tool pprof 进行分析
///debug/pprof/
//
//Types of profiles available:
//Count	Profile
//5	allocs
//0	block
//0	cmdline
//5	goroutine
//5	heap
//0	mutex
//0	profile
//10	threadcreate
//0	trace
//full goroutine stack dump
//Profile Descriptions:
//
//allocs: A sampling of all past memory allocations
//block: Stack traces that led to blocking on synchronization primitives
//cmdline: The command line invocation of the current program
//goroutine: Stack traces of all current goroutines
//heap: A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
//mutex: Stack traces of holders of contended mutexes
//profile: CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
//threadcreate: Stack traces that led to the creation of new OS threads
//trace: A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.

func main() {
	arr := make([]string, 1, 1000000000)
	arr[0] = "a"
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			log.Println(1)
		}
	}()
	http.ListenAndServe(":39090", nil)
}
