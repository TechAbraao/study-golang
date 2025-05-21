package utils

// Funções que começam com minúsculo são funções privadas
func sum(x float32, y float32) float32 {
	return (x + y)
}
func div(x float32, y float32) float32 {
	return (x / y)
}

func sub(x float32, y float32) float32 {
	return (x - y)
}

// Funções que começam com maiúsculo são funções públicas
func Calculator(operation string, x float32, y float32) float32 {
	switch operation {
	case "sum":
		return sum(x, y)
	case "div":
		return div(x, y)
	case "sub":
		return sub(x, y)
	default:
		return 0
	}
}
