package main

import (
	"fmt"
	"github.com/sigurn/crc16"
	"strconv"
	"time"
)

func main() {

	datas := make([]byte, 0)
	//1:数据头固定
	dataHeader := byte(0xFE)
	//2:数据体
	in := make([]byte, 0)
	//3:密钥Key
	key := byte(0x00)
	//4:CMD
	cmd := byte(0x11)

	str := "yOTmK50z"
	in = append(in, key)
	in = append(in, cmd)
	in = append(in, uint8(len([]byte(str))))
	in = append(in, []byte(str)...)

	intt := make([]string, 0)
	for _, v := range strByXOR(in) {
		hex1 := strconv.FormatInt(int64(v), 16)
		intt = append(intt, hex1)
	}

	t := time.Now().Format("0102150405")
	fmt.Println(t)

	//fmt.Printf("%#X\n", byte(t))

	datas = append(datas, dataHeader)
	datas = append(datas, byte(0x66))
	datas = append(datas, strByXOR(in)...)
	fmt.Println(datas)

	table := crc16.MakeTable(crc16.CRC16_MODBUS)

	crc := crc16.Checksum(datas, table)
	c := fmt.Sprintf("%X", crc)
	fmt.Println("CRC16==", c)

}
func strByXOR(input []byte) []uint8 {

	res := make([]uint8, 0)
	for i := 0; i < len(input); i++ {
		//s, _ := fmt.Printf("%#X\n", byte(0x34)^input[i])
		res = append(res, byte(0x34)^input[i])
		//rr = append(rr, fmt.Printf("%#X\n", byte(0x34)^input[i]))
	}
	return res

}

func newNum(input uint8) int {
	res, _ := fmt.Printf("%#X\n", byte(0x34)^input)
	return res
}
