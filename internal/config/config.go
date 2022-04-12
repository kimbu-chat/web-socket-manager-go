package config

func Init() {
	initGRPCCleint()
	initValidations()
}

func Close() {
	if err := closeGRPCConnection(); err != nil {
		panic(err)
	}
}
