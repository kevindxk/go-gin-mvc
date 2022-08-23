/**
 * @Author: dexukong
 * @Description:
 * @File:  jwt
 * @Version: 1.0.0
 * @Date: 2022/08/18 10:36
 */

package utils

var jwtSecret = []byte("kkkkdada")

type Clains struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	//jwt.StandardClaims
}
