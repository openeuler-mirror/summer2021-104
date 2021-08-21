package handler

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte("k8ZEH9EpufjJFBma")

type HmacUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}
type MyClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokeString := c.GetHeader("token")
		if tokeString == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "Please Request Token",
			})
			c.Abort()
			return
		}
		token, claims, err := ParseToken(tokeString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "Token Read Error",
			})
			c.Abort()
			return
		}
		c.Set("userId", claims.UserId)
		c.Set("userName", claims.Username)
		c.Next()
	}
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	return token, claims, err
}

func sha512Str(src string) string {
	h := sha512.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func Login(c *gin.Context) {
	var loginStruct LoginStruct
	err := c.BindJSON(&loginStruct)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "Read Data Error",
		})
		return
	}
	fmt.Println("Login Args", loginStruct)
	userInfo := HmacUser{
		Id:       "1",
		Username: loginStruct.Username,
	}
	token, err := GenerateToken(userInfo)
	if err != nil {
		fmt.Println(err, "===")
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "Gen Token Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Login Success",
		"token":   token,
	})
}

func GenerateToken(u HmacUser) (string, error) {
	expirationTime := time.Now().Add(365 * 24 * time.Hour)
	claims := &MyClaims{
		UserId:   u.Id,
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "token",
			Issuer:    "devops",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	CheckError(err)
	return tokenString, nil
}

func Token2User(c *gin.Context) {
	var username string
	if userName, exists := c.Get("userName"); exists {
		log.Println("Username: ", userName)
		username = userName.(string)
	} else {
		c.JSON(200, gin.H{
			"message":      "Please Login",
			"project_list": "",
		})
		return
	}
	var user_realname string
	err := DB.QueryRow("select `user_realname` from `user` where `user_username`=?", username).Scan(&user_realname)
	CheckError(err)
	c.JSON(200, gin.H{
		"message":       "OK",
		"user_realname": user_realname,
	})
}

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := sha512Str(c.PostForm("password"))
	fmt.Println(username)
	fmt.Println(password)
	var user_realname string
	err := DB.QueryRow("select `user_realname` from `user` where `user_username`=? and `user_password`=?", username, password).Scan(&user_realname)
	if err != nil {
		log.Println(err.Error())
	}
	var succ bool
	var stmt string
	if user_realname == "" {
		succ = false
		stmt = "用户名或密码错误！"
		c.JSON(http.StatusOK, gin.H{
			"code":          0,
			"message":       "Login Success",
			"token":         "",
			"succ":          succ,
			"stmt":          stmt,
			"user_realname": "",
		})
	} else {
		succ = true
		stmt = "OK!"
		userInfo := HmacUser{
			Id:       "1",
			Username: username,
		}
		token, err := GenerateToken(userInfo)
		if err != nil {
			fmt.Println(err, "===")
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "Gen Token Error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":          0,
			"message":       "Login Success",
			"token":         token,
			"succ":          succ,
			"stmt":          stmt,
			"user_realname": user_realname,
		})
	}
}
