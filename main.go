package main
/*
#include "cpuid_rdtsc.h"
*/
import "C"
import (
    "fmt"
    "sort"
    "unsafe"
)

func cpuidCycleCountAvg(low, samples, high int) uint64 {
    var cycles []uint64

    for i := 0; i < low+samples+high; i++ {
        tsc1 := C.read_tsc()
        var eax, ebx, ecx, edx uint32
        C.perform_cpuid((*C.uint)(unsafe.Pointer(&eax)), (*C.uint)(unsafe.Pointer(&ebx)), (*C.uint)(unsafe.Pointer(&ecx)), (*C.uint)(unsafe.Pointer(&edx)))
        tsc2 := C.read_tsc()
        cycles = append(cycles, uint64(tsc2-tsc1))
    }

    // Remove low and high outliers, keep samples
    sort.Slice(cycles, func(i, j int) bool { return cycles[i] < cycles[j] })
    cyclesWithoutOutliers := cycles[low : low+samples]

    // Compute average cycle count without outliers
    var sum uint64
    for _, cycle := range cyclesWithoutOutliers {
        sum += cycle
    }

    return sum / uint64(len(cyclesWithoutOutliers))
}

func insideVMCustom(low, samples, high int, threshold uint64) bool {
    return cpuidCycleCountAvg(low, samples, high) > threshold
}

func insideVM() bool {
    return insideVMCustom(5, 100, 5, 1000)
}

func main() {
    fmt.Println("Inside VM:", insideVM())
	fmt.Scanln()
}
