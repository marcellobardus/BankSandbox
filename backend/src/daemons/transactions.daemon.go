package daemons

import (
	"fmt"
	"time"
)

func TransactionsDaemon() {
	for true {
		time.Sleep(1000 * time.Second)
		fmt.Println("Not implemented yet")
	}
}
