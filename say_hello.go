package gosayhello

func SayHello(name string) string {
	if name == "" {
		return "Hi"
	} else {
		return "Hello " + name + ",\nHave a nice day!"
	}
}
