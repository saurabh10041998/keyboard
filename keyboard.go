package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	fmt.Println("[*] Initiating the github float parser....")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
