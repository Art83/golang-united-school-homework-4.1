package reverse_string

import (
	"bufio"
	"os"
)

func ReverseString(input string) (output string) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	byte_arr := []rune(input)
	for i, j := 0, len(byte_arr)-1; i < j; i, j = i+1, j-1 {
		byte_arr[i], byte_arr[j] = byte_arr[j], byte_arr[i]
	}
	return string(byte_arr)
}
