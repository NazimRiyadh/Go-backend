package standard_named

import "fmt"

//if a func have name then it's standard function
func add(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	p, q := 10, 12
	ans := add(p, q)
	fmt.Println(ans)
}
