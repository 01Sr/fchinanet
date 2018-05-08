/*
* @Author: 01sr
* @Date:   2018-05-08 19:08:10
* @Last Modified by:   01sr
* @Last Modified time: 2018-05-08 19:17:29
 */
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	md5Ctx := md5.New()
	s:="mobile=17751776505&model=01sr-mac&server_did=04741c23-6bd2-42bc-a400-655abf2115e5&time=1525780731349&type=1"
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Print(cipherStr)
	fmt.Print("\n")
	fmt.Print(hex.EncodeToString(cipherStr) + "\n")
	now := time.Now()
	fmt.Println(now.UnixNano() / 1000000)
}
