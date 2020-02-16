package cron

import (
	col "github.com/NJUPT-ISL/Latte/pkg/collection"
	"github.com/NJUPT-ISL/Latte/pkg/log"
	"github.com/NJUPT-ISL/Latte/pkg/ops"
	"github.com/robfig/cron"
	"os"
	"strconv"
	"sync"
)

func UpdateProcessJob(c *cron.Cron) {
	if err := c.AddFunc("0 */5 * * * ?", func() {
		col.MemroyLimit = StrToUInt64(os.Getenv("MEMLIM"))
		if err := col.UpdateProcess();err != nil{
			log.ErrPrint(err)
		}
	}); err != nil {
		log.ErrPrint(err)
	}
}

func UpdateCheckMemJob(c *cron.Cron) {
	if err := c.AddFunc("0 */5 * * * ?", func() {
		if !col.CheckProcess(col.MemroyLimit){
			if err := ops.ReadyToKill(); err != nil{
				log.ErrPrint(err)
			}
		}
	}); err != nil {
		log.ErrPrint(err)
	}
}


func StartJob(c *cron.Cron, w *sync.WaitGroup) {
	w.Add(1)
	c.Start()
}

func StrToUInt64(str string) uint64 {
	if i, e := strconv.ParseUint(str,10,64); e != nil {
		return 0
	} else {
		return i
	}
}
