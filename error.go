package go_helper

func FailOnError(err interface{})  {
	if err != nil {
		panic(err)
	}
}

func CatchError(handler func(err interface{}))  {
	err := recover()
	if err != nil {
		handler(err)
	}
}

func TryCatch(try func(), catch func(err interface{}), finally func())  {
	defer finally()
	defer CatchError(catch)
	try()
}