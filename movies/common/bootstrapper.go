package common

// StartUp loads configurations and connects to DB.
func StartUp() {
	initConfig()
	createDBSession()
}
