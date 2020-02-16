package collection

import (
	"github.com/NJUPT-ISL/Latte/pkg/log"
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
)

type Process struct {
	Name string
	PID uint
	MemoryUsed uint64
}

var (
	Processes  []Process
	MemroyLimit uint64
)
func UpdateProcess() error{
	var NProcesses []Process
	err := nvml.Init()
	if err != nil {
		log.ErrPrint(err)
	}
	defer func() {
		if err := nvml.Shutdown(); err != nil {
			log.ErrPrint(err)
		}
	}()
	count, err := nvml.GetDeviceCount()
	if err != nil {
		log.ErrPrint(err)
	}
	for i := uint(0); i < count; i++ {
		device, err := nvml.NewDevice(i)
		if err != nil {
			log.ErrPrint(err)
		}
		pInfo,err := device.GetAllRunningProcesses()
		if err != nil {
			log.ErrPrint(err)
		}
		for j := range pInfo {
			NProcesses = append(NProcesses,Process{
				Name:       pInfo[j].Name,
				PID:        pInfo[j].PID,
				MemoryUsed: pInfo[j].MemoryUsed,
			})
		}
	}
	Processes = NProcesses
	return err
}

func CheckProcess(mem uint64) bool{
	var sum uint64
	for _,proc := range Processes{
		sum += proc.MemoryUsed
	}
	return sum < mem
}

func GetMaxUsedMemPID() uint{
	var (
		max uint64 = 0
		pid uint = 0
	)
	for _,proc := range Processes{
		if proc.MemoryUsed > max{
			max = proc.MemoryUsed
			pid = proc.PID
		}
	}
	return pid
}

