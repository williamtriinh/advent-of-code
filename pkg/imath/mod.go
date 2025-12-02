package imath

// https://stackoverflow.com/a/59299881
func Mod(a, b int) int {
	return (a%b + b) % b
}
