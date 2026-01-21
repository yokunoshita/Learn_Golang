package helper

// in go capital's matter
// fist letter is uppercase then it can be accessed from another package (like public in other languages)
// first letter is lowercase then it cannot be accessed outside the package (like private in other languages)

var version = "1.2.3"           // this one cannot be accessed outside this package
var Application = "Go Language" // this one can be accessed outside this package

func goodBye(name string) string { // this function cannot be accessed outside this package
	return "Good Bye " + name
}

func Jello(name string) string { // this function can be accessed outside this package
	return "Hello " + name
}

func TryOne() string {
	return version
}
