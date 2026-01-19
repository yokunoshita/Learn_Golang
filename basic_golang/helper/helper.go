package helper

// in go capital's matter
// fist letter is uppercase then it can be accessed from another package
// first letter is lowercase then it cannot be accessed outside the package

var version = "1.2.3"           // this one can't be accessed outside this package
var Application = "Go Language" // this one can

func goodBye(name string) string { // this function cannot be accessed outside this package
	return "Good Bye " + name
}

func Jello(name string) string { // this function can
	return "Hello " + name
}

func TryOne() string {
	return version
}
