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

type Setting struct {
	SettingID     int64  `json:"setting_id"`
	HostURL       string `json:"host_url"`
	SSHPublicKey  string `json:"ssh_public_key"`
	SSHPrivateKey string `json:"ssh_private_key"`
}

type UserList struct {
	User []User `json:"user"`
}

type User struct {
	UserID        int64  `json:"user_id"`
	UserName      string `json:"user_username"`
	UserRealName  string `json:"user_realname"`
	RoleName      string `json:"role_name"`
	RoleID        string `json:"role_id"`
	UserPassword  string `json:"user_password"`
	UserStatus    string `json:"user_status"`
	UserLoginTime string `json:"user_login_time"`
}

type GetUserListRequest struct {
	KeyWords string `json:"KeyWords"`
}

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

func GetSettingInfo() Setting {
	var Setting Setting
	SQL := "select id,host_url,ssh_public_key,ssh_private_key from setting where id=1"
	err := DB.QueryRow(SQL).Scan(&Setting.SettingID, &Setting.HostURL, &Setting.SSHPublicKey, &Setting.SSHPrivateKey)
	CheckError(err)
	return Setting
}

func SetSettingInfo(Setting Setting) {
	stmt, err := DB.Prepare("UPDATE `setting` SET `host_url`=?,`ssh_public_key`=?,`ssh_private_key`=? where id=1")
	CheckError(err)
	_, err = stmt.Exec(Setting.HostURL, Setting.SSHPublicKey, Setting.SSHPrivateKey)
	CheckError(err)
}

func UpdateEvent(UserID int64) {
	SQL := fmt.Sprintf("INSERT INTO event (user_id, event) VALUES (%d, '%s')", UserID, "登陆系统。")
	_ = ExecSQL(SQL)
}

func UpdateLoginTime(UserID int64) {
	CurrentTimeStr := time.Now().Format("2006-01-02 15:04:05")
	SQL := fmt.Sprintf("UPDATE `user` SET user_login_time='%s' WHERE user_id=%d", CurrentTimeStr, UserID)
	_ = ExecSQL(SQL)
}

func ExecSQL(SQL string) int64 {
	stmt, err := DB.Prepare(SQL)
	CheckError(err)
	res, err := stmt.Exec()
	CheckError(err)
	InsertId, err := res.LastInsertId()
	CheckError(err)
	return InsertId
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
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
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

func ReadUserList(c *gin.Context) {
	var Request GetUserListRequest
	c.BindJSON(&Request)
	SQL := "select user.user_id,user.user_username,user.user_realname,role.role_name,user.user_status,user.user_login_time from user,role where user.role_id=role.role_id"
	if Request.KeyWords != "" {
		KeyWords := "%" + Request.KeyWords + "%"
		SQL = fmt.Sprintf("%s and (user.user_username like '%s' or user.user_realname like '%s')", SQL, KeyWords, KeyWords)
	}
	SQL = fmt.Sprintf("%s order by user_id", SQL)
	rows, err := DB.Query(SQL)
	CheckError(err)
	var UserList UserList
	for rows.Next() {
		var User User
		if err := rows.Scan(&User.UserID, &User.UserName, &User.UserRealName, &User.RoleName, &User.UserStatus, &User.UserLoginTime); err != nil {
			fmt.Println(err)
		} else {
			if User.UserLoginTime == "2006-01-02 15:04:05" {
				User.UserLoginTime = ""
			}
			UserList.User = append(UserList.User, User)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "OK",
		"UserList": UserList.User,
	})
}

func CreateUser(c *gin.Context) {
	var User User
	c.BindJSON(&User)
	User.UserPassword = sha512Str(User.UserPassword)
	SQL := fmt.Sprintf("INSERT INTO `user` (`user_username`, `user_password`, `user_realname`, `role_id`) values ('%s','%s','%s',%s)", User.UserName, User.UserPassword, User.UserRealName, User.RoleID)
	log.Println(SQL)
	stmt, err := DB.Prepare(SQL)
	CheckError(err)
	res, err := stmt.Exec()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Error",
		})
		return
	}
	InsertId, err := res.LastInsertId()
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message":  "OK",
		"InsertId": InsertId,
	})
}

func DeleteUserByUserID(c *gin.Context) {
	var User User
	c.BindJSON(&User)
	log.Println(User.UserID)
	stmt, err := DB.Prepare("delete from `user` where `user_id`=?")
	CheckError(err)
	_, err = stmt.Exec(User.UserID)
	CheckError(err)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func ResetPasswordByUserID(c *gin.Context) {
	var User User
	c.BindJSON(&User)
	stmt, err := DB.Prepare("UPDATE `user` SET `user_password`=? where user_id=?")
	CheckError(err)
	_, err = stmt.Exec(sha512Str(User.UserPassword), User.UserID)
	CheckError(err)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := sha512Str(c.PostForm("password"))
	fmt.Println(username)
	fmt.Println(password)
	var user_realname string
	var user_id int64
	var role_id int64
	err := DB.QueryRow("select `user_id`,`user_realname`,`role_id` from `user` where `user_username`=? and `user_password`=?", username, password).Scan(&user_id, &user_realname, &role_id)
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
		go UpdateLoginTime(user_id)
		go UpdateEvent(user_id)
		c.JSON(http.StatusOK, gin.H{
			"code":          0,
			"message":       "Login Success",
			"token":         token,
			"succ":          succ,
			"stmt":          stmt,
			"user_realname": user_realname,
			"user_id":       user_id,
			"role_id":       role_id,
		})
	}
}

func GetSetting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"Setting": GetSettingInfo(),
	})
}

func SetSetting(c *gin.Context) {
	var Setting Setting
	c.BindJSON(&Setting)
	SetSettingInfo(Setting)
	c.JSON(200, gin.H{
		"message": "OK",
	})
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
