package main
import (
	"net"
	"strings"
	"strconv"
	"fmt"
)

// Convert uint to net.IP http://www.outofmemory.cn
func inet_ntoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3],bytes[2],bytes[1],bytes[0])
}

// Convert net.IP to int64 ,  http://www.outofmemory.cn
func inet_aton(ipnr net.IP) int64 {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

//二进制转十六进制
func btox(b string) string {
	base, _ := strconv.ParseInt(b, 2, 10)
	return strconv.FormatInt(base, 16)
}

//十六进制转二进制
func xtob(x string) string {
	base, _ := strconv.ParseInt(x, 16, 10)
	return strconv.FormatInt(base, 2)
}

//十进制转二进制
func dtob(x string) string {
	base, _ := strconv.ParseInt(x, 10, 10)
	return strconv.FormatInt(base, 2)
}

func main() {
	var address = net.IPv4(10, 114, 50, 11)

	fmt.Println(address)

	var result = inet_aton(address)
	fmt.Println(result)

	fmt.Println(dtob("10"))
	fmt.Println(dtob("114"))
	fmt.Println(dtob("50"))
	fmt.Println(dtob("11"))
	fmt.Println(dtob("128"))
}