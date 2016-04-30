package common

func StartUp() {
	initConfig()
	initKey()
	createDbSession()
	addIndexes()
}
