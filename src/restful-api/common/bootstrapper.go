package common

import (
	"log"
)

func StartUp() {
	initConfig()
	log.Println("[Common:Bootstrapper:InitConfig] success----->")
	initKeys()
	log.Println("[Common:Bootstrapper:InitKeys] success----->")
	createDbSession()
	log.Println("[Common:Bootstrapper:CreateDbSession] success----->")
	addIndexes()
	log.Println("[Common:Bootstrapper:AddIndexes] success----->")
}
