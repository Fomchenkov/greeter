package greeter

// Greet func
func Greet(name string) string {
	return "Hello " + name + "!"
}

// GreetAny func
func GreetAny(names []string) string {
	res := ""
	for i := 0; i < len(names); i++ {
		res += "Hello " + names[i] + "!\n"
	}
	return res[0:len(res)-1]
}
