package utils

func YourName(name string) string {
	if name == "" {
		return "Undefined name."
	}
	return "Your name is: " + name
}
