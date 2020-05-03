package main

import (
	"log"
)

func main() {
	if err := initInstaClient(); err != nil {
		log.Fatalln("unable to init the instagram client")
	}

	for _, action := range Config.Actions {
		err := performAction(action.Name, action.Args)
		if err != nil {
			log.Printf("could not perform action '%s'. err: %s", action.Name, err.Error())
		}
	}
}
