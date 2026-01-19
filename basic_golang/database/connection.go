package database

var connect string

// if your function would be called as soon as the package is called then just name the function init()
func init() {
	connect = "Mysql - this is an example of func being called from another package as soon as the package is calleds"
}

func GetDatabase() string {
	return connect
}
