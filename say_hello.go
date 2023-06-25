package gosayhello

func SayHello(name string, age int, sex string) string {
	if name == "" {
		return "Hi"
	} else {
		var greeting string
		switch age >= 25 {
		case true:
			if sex == "male" {
				greeting = "Mr."
			} else if sex == "female" {
				greeting = "Ms."
			}
		case false:
			greeting = ""
		}

		return "Hello " + greeting + name + ",\nHave a nice day!"
	}
}
