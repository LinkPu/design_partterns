package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Flag int
}

// hungry
var hungryInstance = &Singleton{}

func HungrySingleton() *Singleton {
	return hungryInstance
}

// lazy
var lazyInstance *Singleton

func LazySingleton() *Singleton {
	if lazyInstance == nil {
		// 当第一次运行到这里时如果有其他协程并发创建实例，lazyInstance 仍为 nil，因此不能防并发
		lazyInstance = &Singleton{}
	}

	return lazyInstance
}

// lazy with lock
var mutex sync.Mutex

func SyncLazySingleton() *Singleton {
	// lock the flow expression, but lock and unlock every time
	mutex.Lock()
	defer mutex.Unlock()

	if lazyInstance == nil {
		lazyInstance = &Singleton{}
	}

	return lazyInstance
}

func SyncLazySingletonOptimize() *Singleton {
	if lazyInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()

		if lazyInstance == nil {
			lazyInstance = &Singleton{}
		}
	}

	return lazyInstance
}

// sync.Once
var once sync.Once

func SyncLazySingletonByOnce() *Singleton {
	once.Do(func() {
		lazyInstance = &Singleton{}
	})

	return lazyInstance
}

func main() {
	// initFunc := HungrySingleton
	// initFunc := LazySingleton
	// initFunc := SyncLazySingleton
	// initFunc := SyncLazySingletonOptimize
	initFunc := SyncLazySingletonByOnce

	instance := initFunc()
	instance.Flag = 1
	fmt.Printf("instance: %++v, ptr: %p\n\n", instance, instance)

	instance1 := initFunc()
	fmt.Printf("instance1: %++v, ptr: %p\n\n", instance1, instance1)

	instance1.Flag = 2
	fmt.Printf("isntance: %++v\n", instance)
	fmt.Printf("instance1: %++v\n\n", instance1)
}
