/**
 * @Author: dexukong
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2022/08/23 10:16
 */

package serializer

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
