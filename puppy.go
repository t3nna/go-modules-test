package golang_test_modules

func Bark() string {
	return "Bark"
}

func Barks() string {
	return Bark() + " " + Bark()
}
