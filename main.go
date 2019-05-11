package main

import "log"

func main() {
	initInstaClient()

	for _, action := range Config.Actions {
		err := performAction(action.Name, action.Args)
		if err != nil {
			log.Printf("could not perform action '%s'. err: %s", action.Name, err.Error())
		}
	}
}
