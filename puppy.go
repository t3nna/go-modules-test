package golang_test_modules

import "github.com/GoesToEleven/dog"

func Bark() string {
	return "Bark"
}

func Barks() string {
	return Bark() + " " + Bark()
}

func MessageFromThirdPackage(s string) string {
	return dog.WhenGrownUp(s)
}
