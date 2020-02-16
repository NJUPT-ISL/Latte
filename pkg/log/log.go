package log

import (
	"fmt"
	"time"
)

func ErrPrint(err error) {
	fmt.Printf("[Latte]-[%s] Error: %s\n", time.Now(), err)
}

func Print(log string) {
	fmt.Printf("[Latte]-[%s]  %s\n", time.Now(), log)
}