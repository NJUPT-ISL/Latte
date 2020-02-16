package ops

import (
	col "github.com/NJUPT-ISL/Latte/pkg/collection"
	"github.com/NJUPT-ISL/Latte/pkg/log"
	"syscall"
	"time"
)

func killProcess(pid uint) error{
	return syscall.Kill(int(pid), syscall.SIGKILL)
}

func ReadyToKill() error{
	time.Sleep(3*time.Minute)
	if err := col.UpdateProcess(); err !=nil{
		log.ErrPrint(err)
		return err
	}
	KillLoop()
	return nil
}

func KillLoop(){
	if !col.CheckProcess(col.MemroyLimit){
		if err := killProcess(col.GetMaxUsedMemPID()); err != nil{
			log.ErrPrint(err)
			return
		}
		if err := col.UpdateProcess();err != nil{
			log.ErrPrint(err)
			return
		}
		KillLoop()
	}
}
