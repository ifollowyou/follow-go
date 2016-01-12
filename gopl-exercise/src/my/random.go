package main
import (
	"crypto/rand"
	"math/big"
	"fmt"
	"strconv"
)

func main() {
	max := big.NewInt(100)
	i, _ := rand.Int(rand.Reader, max)
	s := fmt.Sprintf("%d", i)
	strconv.Atoi(s)
	fmt.Println(s)
}
