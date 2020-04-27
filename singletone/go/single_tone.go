package main

import (
	"fmt"
	"sync"
)

type Singletone struct {
	Flag int
}

// hungry
var hungryInstance = &Singletone{}

func HungrySingletone() *Singletone {
	return hungryInstance
}

// lazy
var lazyInstance *Singletone

func LazySingletone() *Singletone {
	if lazyInstance == nil {
		// 当第一次运行到这里时如果有其他协程并发创建实例，lazyInstance 仍为 nil，因此不能防并发
		lazyInstance = &Singletone{}
	}

	return lazyInstance
}

// lazy with lock
var mutex sync.Mutex

func SyncLazySingletone() *Singletone {
	// lock the flow expression, but lock and unlock every time
	mutex.Lock()
	defer mutex.Unlock()

	if lazyInstance == nil {
		lazyInstance = &Singletone{}
	}

	return lazyInstance
}

func SyncLazySingletoneOptimize() *Singletone {
	if lazyInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()

		if lazyInstance == nil {
			lazyInstance = &Singletone{}
		}
	}

	return lazyInstance
}

// sync.Once
var once sync.Once

func SyncLazySingletoneByOnce() *Singletone {
	once.Do(func() {
		lazyInstance = &Singletone{}
	})

	return lazyInstance
}

func main() {
	// initFunc := HungrySingletone
	// initFunc := LazySingletone
	// initFunc := SyncLazySingletone
	// initFunc := SyncLazySingletoneOptimize
	initFunc := SyncLazySingletoneByOnce

	instance := initFunc()
	instance.Flag = 1
	fmt.Printf("instance: %++v, ptr: %p\n\n", instance, instance)

	instance1 := initFunc()
	fmt.Printf("instance1: %++v, ptr: %p\n\n", instance1, instance1)

	instance1.Flag = 2
	fmt.Printf("isntance: %++v\n", instance)
	fmt.Printf("instance1: %++v\n\n", instance1)
}
