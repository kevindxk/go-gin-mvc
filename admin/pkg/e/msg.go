/**
 * @Author: dexukong
 * @Description:
 * @File:  msg
 * @Version: 1.0.0
 * @Date: 2022/08/23 10:20
 */

package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	} else {
		return msg
	}
}
