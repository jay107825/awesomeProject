package main

/**
 * @Author: admin
 * @Description:  性能
 * @File: profile
 * @Version: 1.0.0
 * @Date: 2023/5/28 21:15
 */

func joinSlice() []string {
	var arr []string

	for i := 0; i < 10000; i++ {
		// 故意造成多次切片添加（add）操作，由于每次操作可能会有内存重新分配和移动
		arr = append(arr, "arr")
	}

	return arr
}

func main() {
	// p.Stop() must be called before the program exits to
	// ensure profiling information is written to disk.
	//p := Start(MemProfile, ProfilePath("."), NoShutdownHook)

	// You can enable different kinds of memory profiling, either Heap or Allocs where Heap
	// profiling is the default with profile.MemProfile.
	//p := Start(MemProfileAllocs, ProfilePath("."), NoShutdownHook)
}
