package core

import (
	"fmt"
	"reflect"
	"time"
)

type queueItem struct {
	Process reflect.Value
	Args    []reflect.Value
	Channel chan bool
}

var queue []queueItem

func AddStatementToQueue(statement string) error {
	return nil
}

func StartQueue() {
	ticker := time.NewTicker(time.Duration(400) * time.Millisecond)
	go func() {
		for range ticker.C {
			processQueue()
		}
	}()
}

func processQueue() {
	if len(queue) == 0 {
		return
	}

	itemToProcess := queue[0]
	functionResponse := itemToProcess.Process.Call(itemToProcess.Args)

	err, _ := functionResponse[0].Interface().(error)
	if err != nil {
		itemToProcess.Channel <- false
		queue = queue[1:]
		fmt.Println(err)
		return
	}

	itemToProcess.Channel <- true

	queue = queue[1:]
}
