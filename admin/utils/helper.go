/**
 * @Author: dexukong
 * @Description:
 * @File:  helper
 * @Version: 1.0.0
 * @Date: 2022/08/08 15:38
 */

package utils

import (
	uuid "github.com/satori/go.uuid"
)

//type UserClaims struct {
//	Identity string `json:"identity"`
//	Name     string `json:"name"`
//	IsAdmin  int    `json:"is_admin"`
//	jwt.StandardClaims
//}
//
//// GetMd5
//// 生成 md5
//func GetMd5(s string) string {
//	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
//}
//
//var myKey = []byte("gin-gorm-oj-key")
//
//// GenerateToken
//// 生成 token
//func GenerateToken(identity, name string, isAdmin int) (string, error) {
//	UserClaim := &UserClaims{
//		Identity:       identity,
//		Name:           name,
//		IsAdmin:        isAdmin,
//		StandardClaims: jwt.StandardClaims{},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
//	tokenString, err := token.SignedString(myKey)
//	if err != nil {
//		return "", err
//	}
//	return tokenString, nil
//}
//
//// AnalyseToken
//// 解析 token
//func AnalyseToken(tokenString string) (*UserClaims, error) {
//	userClaim := new(UserClaims)
//	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
//		return myKey, nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if !claims.Valid {
//		return nil, fmt.Errorf("analyse Token Error:%v", err)
//	}
//	return userClaim, nil
//}
//
//// SendCode
//// 发送验证码
//func SendCode(toUserEmail, code string) error {
//	e := email.NewEmail()
//	e.From = "Get <getcharzhaopan@163.com>"
//	e.To = []string{toUserEmail}
//	e.Subject = "验证码已发送，请查收"
//	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
//	return e.SendWithTLS("smtp.163.com:465",
//		smtp.PlainAuth("", "18100891626@163.com", define.MailPassword, "smtp.163.com"),
//		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
//}
//
// GetUUID
// 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
}

//
//// GetRand
//// 生成验证码
//func GetRand() string {
//	rand.Seed(time.Now().UnixNano())
//	s := ""
//	for i := 0; i < 6; i++ {
//		s += strconv.Itoa(rand.Intn(10))
//	}
//	return s
//}
//
//// CodeSave
//// 保存代码
//func CodeSave(code []byte) (string, error) {
//	dirName := "code/" + GetUUID()
//	path := dirName + "/main.go"
//	err := os.Mkdir(dirName, 0777)
//	if err != nil {
//		return "", err
//	}
//	f, err := os.Create(path)
//	if err != nil {
//		return "", err
//	}
//	f.Write(code)
//	defer f.Close()
//	return path, nil
//}
