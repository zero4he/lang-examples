package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func main() {

	sum := md5.Sum([]byte("a"))
	fmt.Println(sum)
	fmt.Printf("%s\n", hex.EncodeToString(sum[:])) // 输出加密结果
	rv := ((sum[3] & 0xFF) << 24) | ((sum[2] & 0xFF) << 16) | ((sum[1] & 0xFF) << 8) | (sum[0] & 0xFF)
	fmt.Println(rv)
	hash := int64(rv) & 0xffffffff;
	fmt.Println(hash)

}
