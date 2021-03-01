package memory

import (
	"sync"
	"syscall"
)

type Memory_t struct {
	Free uint64
	Total uint64
	SwapFree uint64
	SwapTotal uint64
	mu sync.Mutex
}

func Get() (*Memory_t, error) {
	sysInfo := &syscall.Sysinfo_t{}
	memoryInfo := &Memory_t{}

	err := syscall.Sysinfo(sysInfo)
	if err != nil {
		return nil, err
	}

	defer memoryInfo.mu.Unlock()

	memoryInfo.mu.Lock()
	memoryInfo.Free = uint64(sysInfo.Freeram)
	memoryInfo.Total = uint64(sysInfo.Totalram)
	memoryInfo.SwapFree = uint64(sysInfo.Freeswap)
	memoryInfo.SwapTotal = uint64(sysInfo.Totalswap)

	return memoryInfo, nil
}
