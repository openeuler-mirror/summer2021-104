package handler

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

var (
	ConfigContent Config
	DB            *sql.DB
)

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type MonitorRequest struct {
	HostID string `json:"host_id"`
}

type MonitorData struct {
	Value float32 `json:"value"`
	Time  string  `json:"time"`
}

type LogStruct struct {
	Username  string `json:"username"`
	Event     string `json:"event"`
	EventTime string `json:"event_time"`
}

type Config struct {
	Database Database `yaml:"database"`
}

type Role struct {
	RoleID     int64  `json:"role_id"`
	RoleName   string `json:"role_name"`
	UserNum    int64  `json:"user_num"`
	RoleRemark string `json:"role_remark"`
}

func MonitorDAO(HostID string, KeyType string, Value float64) {
	SQL := fmt.Sprintf("INSERT INTO devops.monitor_metric (host_id, `key`, value) VALUES(%s, '%s', %f)", HostID, KeyType, Value)
	stmt, err := DB.Prepare(SQL)
	CheckError(err)
	_, err = stmt.Exec()
	CheckError(err)
}

func Monitor() {
	HostList := GetHostList("").HostManagement
	for i := 0; i < len(HostList); i++ {
		CPUUsed := ExecCmdFunc(HostList[i].HostManagementID, "top -b -n 1 | grep Cpu | awk '{print $8}' | cut -f 1 -d '%'")
		MemoryTotal := ExecCmdFunc(HostList[i].HostManagementID, "free | grep Mem | awk '{print $2}'")
		MemoryAvailable := ExecCmdFunc(HostList[i].HostManagementID, "free | grep Mem | awk '{print $7}'")
		CPUUsedFloat, err := strconv.ParseFloat(CleanString(CPUUsed), 32)
		if err == nil {
			value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 100-CPUUsedFloat), 64)
			MonitorDAO(HostList[i].HostManagementID, "CPU", value)
		}
		MemoryTotalInt, err1 := strconv.Atoi(CleanString(MemoryTotal))
		MemoryAvailableInt, err2 := strconv.Atoi(CleanString(MemoryAvailable))
		if err1 == nil && err2 == nil {
			MemortUsage := (1 - (float32(MemoryAvailableInt) / float32(MemoryTotalInt))) * 100
			value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", MemortUsage), 64)
			MonitorDAO(HostList[i].HostManagementID, "Memory", value)
		}
	}
}

func MonitorCircle() {
	for {
		Monitor()
		time.Sleep(20 * time.Second)
	}
}

func CleanString(S string) string {
	S = strings.Replace(S, " ", "", -1)
	D := strings.Replace(S, "\n", "", -1)
	return D
}

func DeleteRoleByRoleID(c *gin.Context) {
	var Role Role
	c.BindJSON(&Role)
	stmt, err := DB.Prepare("delete from `role` where `role_id`=?")
	CheckError(err)
	_, err = stmt.Exec(Role.RoleID)
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func GetCPU(c *gin.Context) {
	var Request MonitorRequest
	c.BindJSON(&Request)
	var Time []string
	var Value []float32
	SQL := fmt.Sprintf("SELECT value, create_time FROM monitor_metric WHERE host_id=%s and `key`='CPU' order by create_time limit 100", Request.HostID)
	rows, err := DB.Query(SQL)
	CheckError(err)
	for rows.Next() {
		var MonitorData MonitorData
		if err := rows.Scan(&MonitorData.Value, &MonitorData.Time); err != nil {
			fmt.Println(err)
		} else {
			Time = append(Time, MonitorData.Time)
			Value = append(Value, MonitorData.Value)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"Time":    Time,
		"Value":   Value,
	})
}

func GetMemory(c *gin.Context) {
	var Request MonitorRequest
	c.BindJSON(&Request)
	var Time []string
	var Value []float32
	SQL := fmt.Sprintf("SELECT value, create_time FROM monitor_metric WHERE host_id=%s and `key`='Memory' order by create_time limit 100", Request.HostID)
	rows, err := DB.Query(SQL)
	CheckError(err)
	for rows.Next() {
		var MonitorData MonitorData
		if err := rows.Scan(&MonitorData.Value, &MonitorData.Time); err != nil {
			fmt.Println(err)
		} else {
			Time = append(Time, MonitorData.Time)
			Value = append(Value, MonitorData.Value)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"Time":    Time,
		"Value":   Value,
	})
}

func GetLog(c *gin.Context) {
	var LogStructList []LogStruct
	SQL := "SELECT `user`.`user_realname`, event, event_create_time FROM `event`,`user` where `user`.`user_id`=`event`.`user_id`"
	rows, err := DB.Query(SQL)
	CheckError(err)
	for rows.Next() {
		var LogStruct LogStruct
		if err := rows.Scan(&LogStruct.Username, &LogStruct.Event, &LogStruct.EventTime); err != nil {
			fmt.Println(err)
		} else {
			LogStructList = append(LogStructList, LogStruct)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message":    "OK",
		"event_list": LogStructList,
	})
}

func ReadRoleList(c *gin.Context) {
	var RoleList []Role
	SQL := "select r.role_id, r.role_name,  IFNULL(u.ct, 0) as user_num, r.role_remark from role r left join (select role_id, count(1) as ct from user group by role_id) as u on u.role_id=r.role_id"
	rows, err := DB.Query(SQL)
	CheckError(err)
	for rows.Next() {
		var Role Role
		if err := rows.Scan(&Role.RoleID, &Role.RoleName, &Role.UserNum, &Role.RoleRemark); err != nil {
			fmt.Println(err)
		} else {
			RoleList = append(RoleList, Role)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "OK",
		"RoleList": RoleList,
	})
}

func CreateRole(c *gin.Context) {
	var Role Role
	c.BindJSON(&Role)
	SQL := fmt.Sprintf("INSERT INTO `role` (`role_name`, `role_remark`) values ('%s','%s')", Role.RoleName, Role.RoleRemark)
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

func CheckError(err error) {
	if err == sql.ErrNoRows {
		log.Println(err)
	} else if err != nil {
		panic(err)
	}
}

func (c *Config) GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("config/config.yml")
	CheckError(err)
	err = yaml.Unmarshal(yamlFile, c)
	CheckError(err)
	return c
}

func InitConfig() {
	ConfigContent.GetConf()
}

func InitDB(database Database) {
	username := database.Username
	password := database.Password
	host := database.Host
	port := database.Port
	database1 := database.Database
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, database1)
	fmt.Println(conn)
	DB, _ = sql.Open("mysql", conn)
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	err := DB.Ping()
	if err != nil {
		panic(err)
	}
}

func Healthz(c *gin.Context) {

}

type HostCount struct {
	Active   int `json:"active"`
	Inactive int `json:"inactive"`
}

type UserCount struct {
	System int `json:"system"`
	Ops    int `json:"ops"`
	Dev    int `json:"dev"`
	Test   int `json:"test"`
}

type DashboardInfoResponse struct {
	Host HostCount `json:"host"`
	User UserCount `json:"user"`
}

func DashboardInfo(c *gin.Context) {
	var Response DashboardInfoResponse
	SQL := "select count(*) as active from host_management WHERE host_status=1"
	err := DB.QueryRow(SQL).Scan(&Response.Host.Active)
	CheckError(err)
	SQL = "select count(*) as inactive from host_management WHERE host_status=0"
	err = DB.QueryRow(SQL).Scan(&Response.Host.Inactive)
	CheckError(err)
	SQL = "select count(*) as system from user,role where user.role_id=role.role_id and role.role_name='系统管理'"
	err = DB.QueryRow(SQL).Scan(&Response.User.System)
	CheckError(err)
	SQL = "select count(*) as ops from user,role where user.role_id=role.role_id and role.role_name='运维团队'"
	err = DB.QueryRow(SQL).Scan(&Response.User.Ops)
	CheckError(err)
	SQL = "select count(*) as dev from user,role where user.role_id=role.role_id and role.role_name='开发团队'"
	err = DB.QueryRow(SQL).Scan(&Response.User.Dev)
	CheckError(err)
	SQL = "select count(*) as test from user,role where user.role_id=role.role_id and role.role_name='测试团队'"
	err = DB.QueryRow(SQL).Scan(&Response.User.Test)
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message":  "OK",
		"Response": Response,
	})
}

func PublicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

func ExecCmdFunc(HostID string, CMD string) string {
	Host := SelectHostByHostID(HostID)
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            Host.HostManagementUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	Settings := GetSettingInfo()
	signer, err := ssh.ParsePrivateKey([]byte(Settings.SSHPrivateKey))
	config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	addr := fmt.Sprintf("%s:%s", Host.HostManagementIP, Host.HostManagementPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	CheckError(err)
	defer sshClient.Close()
	session, err := sshClient.NewSession()
	CheckError(err)
	defer session.Close()
	combo, err := session.CombinedOutput(CMD)
	CheckError(err)
	return string(combo)
}

type ExecCmdRequest struct {
	HostID string `json:"host_id"`
	CMD    string `json:"cmd"`
}

func ExecCmd(c *gin.Context) {
	var Request ExecCmdRequest
	c.BindJSON(&Request)
	stdout := ExecCmdFunc(Request.HostID, Request.CMD)
	c.JSON(200, gin.H{
		"message": "OK",
		"stdout":  stdout,
	})
}

type HostManagement struct {
	HostManagementID            string `json:"host_management_id"`
	HostManagementIP            string `json:"host_management_ip"`
	HostManagementName          string `json:"host_management_name"`
	HostManagementUser          string `json:"host_management_user"`
	HostManagementPassword      string `json:"host_management_password"`
	HostManagementKey           string `json:"host_management_key"`
	HostManagementConnectMethod string `json:"host_management_connect_method"`
	HostManagementCreateTime    string `json:"host_management_create_time"`
	HostManagementModifyTime    string `json:"host_management_modify_time"`
	HostManagementPort          string `json:"host_management_port"`
}

type HostManagementList struct {
	HostManagement []HostManagement `json:"host_management"`
}

type CreateHostRequest struct {
	HostIP       string `json:"HostIP"`
	HostPort     string `json:"HostPort"`
	HostName     string `json:"HostName"`
	HostUser     string `json:"HostUser"`
	HostPassword string `json:"HostPassword"`
}

func CreateHost(c *gin.Context) {
	var Request CreateHostRequest
	c.BindJSON(&Request)
	err := HostPasswordCheck(Request.HostIP, Request.HostPort, Request.HostUser, Request.HostPassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户名或密码错误",
		})
		return
	}
	SQL := "INSERT INTO `host_management` (`host_port`,`host_management_ip`,`host_management_name`,`host_management_user`)"
	SQL = fmt.Sprintf("%s values ('%s','%s','%s','%s')", SQL, Request.HostPort, Request.HostIP, Request.HostName, Request.HostUser)
	stmt, err := DB.Prepare(SQL)
	CheckError(err)
	res, err := stmt.Exec()
	InsertId, err := res.LastInsertId()
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message":   "OK",
		"insert_id": InsertId,
	})
}

func UpdateHost(c *gin.Context) {
	HostID := c.PostForm("HostID")
	HostIP := c.PostForm("HostIP")
	HostName := c.PostForm("HostName")
	HostUser := c.PostForm("HostUser")
	HostPassword := c.PostForm("HostPassword")
	HostKey := c.PostForm("HostKey")
	HostConnectMethod := c.PostForm("HostConnectMethod")
	stmt, err := DB.Prepare("UPDATE `host_management` SET `host_management_ip`=?,`host_management_name`=?,`host_management_user`=?,`host_management_password`=?,`host_management_key`=?,`host_management_connect_method`=? where host_management_id=?")
	CheckError(err)
	_, err = stmt.Exec(HostIP, HostName, HostUser, HostPassword, HostKey, HostConnectMethod, HostID)
	CheckError(err)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

type GetHostListRequest struct {
	KeyWords string `json:"KeyWords"`
}

func GetHostList(KeyWord string) HostManagementList {
	var host_management_list HostManagementList
	SQL := "select `host_port`,`host_management_id`,`host_management_ip`,`host_management_name`,`host_management_user`,`host_management_create_time`,`host_management_modify_time` from `host_management`"
	if KeyWord != "" {
		KeyWords := "%" + KeyWord + "%"
		SQL = fmt.Sprintf("%s  where host_management_ip like '%s' or host_management_name like '%s'", SQL, KeyWords, KeyWords)
	}
	rows, err := DB.Query(SQL)
	CheckError(err)
	for rows.Next() {
		var HostPort, host_management_id, host_management_ip, host_management_name, host_management_user, host_management_create_time, host_management_modify_time string
		if err := rows.Scan(&HostPort, &host_management_id, &host_management_ip, &host_management_name, &host_management_user, &host_management_create_time, &host_management_modify_time); err != nil {
			fmt.Println(err)
		} else {
			host_management_ip = fmt.Sprintf("%s:%s", host_management_ip, HostPort)
			host_management_list.HostManagement = append(host_management_list.HostManagement,
				HostManagement{HostManagementID: host_management_id, HostManagementIP: host_management_ip, HostManagementName: host_management_name, HostManagementUser: host_management_user, HostManagementCreateTime: host_management_create_time, HostManagementModifyTime: host_management_modify_time})
		}
	}
	return host_management_list
}

func ReadHostList(c *gin.Context) {
	var Request GetHostListRequest
	c.BindJSON(&Request)
	c.JSON(200, gin.H{
		"message":              "OK",
		"host_management_list": GetHostList(Request.KeyWords),
	})
}

func SelectHostByHostID(HostID string) HostManagement {
	var host_management HostManagement
	err := DB.QueryRow("select `host_port`,`host_management_ip`,`host_management_user` from `host_management` where `host_management_id`=?", HostID).Scan(&host_management.HostManagementPort, &host_management.HostManagementIP, &host_management.HostManagementUser)
	CheckError(err)
	return host_management
}

func ReadHost(c *gin.Context) {
	HostID := c.PostForm("HostID")
	c.JSON(200, gin.H{
		"message":         "OK",
		"host_management": SelectHostByHostID(HostID),
	})
}

func DeleteHost(c *gin.Context) {
	HostID := c.PostForm("HostID")
	SQL := fmt.Sprintf("delete from `host_management` where `host_management_id`=%s", HostID)
	stmt, err := DB.Prepare(SQL)
	CheckError(err)
	_, err = stmt.Exec()
	CheckError(err)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func HostPasswordCheck(HostAddress string, HostPort string, HostUser string, Password string) error {
	log.Println("HostPort: ", HostPort)
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            HostUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(Password)}
	addr := fmt.Sprintf("%s:%s", HostAddress, HostPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Println("Create SSH Client Failure.", err)
		return err
	}
	defer sshClient.Close()
	session, err := sshClient.NewSession()
	if err != nil {
		log.Println("Create SSH Session Failure.", err)
		return err
	}
	defer session.Close()
	SystemPublicKey := GetSettingInfo().SSHPublicKey
	CMD := fmt.Sprintf("echo \"%s\" >> .ssh/authorized_keys", SystemPublicKey)
	combo, err := session.CombinedOutput(CMD)
	if err != nil {
		log.Fatal("Remote Exec CMD Failure.", err)
	}
	log.Println("CMD Output:", string(combo))
	return nil
}
