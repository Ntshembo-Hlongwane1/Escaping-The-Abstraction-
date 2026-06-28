package main

import (
	"fmt"

	"searchengine.com/src/internal/engine"
	"searchengine.com/src/internal/kernal"
	"searchengine.com/src/internal/store"
)

func main() {
	fmt.Printf("\n Search Engine")

	k := kernal.NewKernal()

	fmt.Printf("\n System Kernal Initializing")

	store := store.NewStore()
	fileReader := kernal.NewFileReader("./data", store)
	engine := engine.NewEngine(store)

	k.Register(fileReader)
	k.Register(store)
	k.Register(engine)

	if err := k.InitAll(); err != nil {
		fmt.Printf("\n Fatal %v\n", err)
		k.Shutdown()
		return
	}

	if err := k.StartAll(); err != nil {
		fmt.Printf("Fatal %v \n", err)
		k.Shutdown()
		return
	}

	fmt.Printf("\n System is running")
}
