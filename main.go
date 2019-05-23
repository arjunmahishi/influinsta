package main

import (
	"log"
	"sync"
)

func main() {
	initInstaClient()

	var wg sync.WaitGroup
	for _, action := range Config.Actions {
		wg.Add(1)
		go func(name string, args []string) {
			err := performAction(name, args)
			if err != nil {
				log.Printf("could not perform action '%s'. err: %s", name, err.Error())
			}
			wg.Done()
		}(action.Name, action.Args)
	}
	wg.Wait()
}
