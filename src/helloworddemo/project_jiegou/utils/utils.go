package utils

//运算符计算.
func OperFunc(value1 float64, value2 float64, oper byte) (result float64) {

	switch oper {
	case '+':
		result = value1 + value2
	case '-':
		result = value1 - value2
	case '*':
		result = value1 * value2
	case '/':
		result = value1 / value2
	default:
		result = 0.0
	}

	return result
}
