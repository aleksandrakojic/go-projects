package main

import (
    "fmt"
    "sync"
)

var lock = &sync.Mutex{}

type single struct {

}
var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
			fmt.Println("Creating single instance")
            singleInstance = &single{}
        } else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Reusing single instance")
	}
	return singleInstance
}