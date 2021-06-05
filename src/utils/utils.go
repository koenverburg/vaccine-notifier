package utils

func CheckIfErr(e error) {
	if e != nil {
		panic(e)
	}
}
