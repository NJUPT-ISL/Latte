package main

import (
	"github.com/NJUPT-ISL/Latte/pkg/collection"
	job "github.com/NJUPT-ISL/Latte/pkg/cron"
	"github.com/NJUPT-ISL/Latte/pkg/log"
	"github.com/robfig/cron"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	collection.MemroyLimit = job.StrToUInt64(os.Getenv("MEMLIM"))
	if err := collection.UpdateProcess(); err != nil{
		log.ErrPrint(err)
	}
	c := cron.New()
	job.UpdateProcessJob(c)
	job.UpdateCheckMemJob(c)
	defer c.Stop()
	job.StartJob(c, &wg)
	wg.Wait()
}