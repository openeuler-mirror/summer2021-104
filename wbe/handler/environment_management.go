package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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
}

type HostManagementList struct {
	HostManagement []HostManagement `json:"host_management"`
}

func CreateHost(c *gin.Context) {
	HostIP := c.PostForm("HostIP")
	HostName := c.PostForm("HostName")
	HostUser := c.PostForm("HostUser")
	HostPassword := c.PostForm("HostPassword")
	HostKey := c.PostForm("HostKey")
	HostConnectMethod := c.PostForm("HostConnectMethod")
	stmt, err := DB.Prepare("INSERT INTO `host_management` (`host_management_ip`,`host_management_name`,`host_management_user`,`host_management_password`,`host_management_key`,`host_management_connect_method`) values (?,?,?,?,?,?)")
	CheckError(err)
	res, err := stmt.Exec(HostIP, HostName, HostUser, HostPassword, HostKey, HostConnectMethod)
	InsertId, err := res.LastInsertId()
	CheckError(err)
	c.JSON(200, gin.H{
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
func ReadHostList(c *gin.Context) {
	var host_management_list HostManagementList
	rows, err := DB.Query("select `host_management_id`,`host_management_ip`,`host_management_name`,`host_management_user`,`host_management_password`,`host_management_key`,`host_management_connect_method`,`host_management_create_time`,`host_management_modify_time` from `host_management`")
	CheckError(err)
	for rows.Next() {
		var host_management_id, host_management_ip, host_management_name, host_management_user, host_management_password, host_management_key, host_management_connect_method, host_management_create_time, host_management_modify_time string
		if err := rows.Scan(&host_management_id, &host_management_ip, &host_management_name, &host_management_user, &host_management_password, &host_management_key, &host_management_connect_method, &host_management_create_time, &host_management_modify_time); err != nil {
			fmt.Println(err)
		} else {
			host_management_list.HostManagement = append(host_management_list.HostManagement,
				HostManagement{HostManagementID: host_management_id, HostManagementIP: host_management_ip, HostManagementName: host_management_name, HostManagementUser: host_management_user, HostManagementPassword: host_management_password, HostManagementKey: host_management_key, HostManagementConnectMethod: host_management_connect_method, HostManagementCreateTime: host_management_create_time, HostManagementModifyTime: host_management_modify_time})
		}
	}
	c.JSON(200, gin.H{
		"message":              "OK",
		"host_management_list": host_management_list,
	})
}

func ReadHost(c *gin.Context) {
	var host_management HostManagement
	HostID := c.PostForm("HostID")
	var host_management_id, host_management_ip, host_management_name, host_management_user, host_management_password, host_management_key, host_management_connect_method, host_management_create_time, host_management_modify_time string
	err := DB.QueryRow("select `host_management_id`,`host_management_ip`,`host_management_name`,`host_management_user`,`host_management_password`,`host_management_key`,`host_management_connect_method`,`host_management_create_time`,`host_management_modify_time` from `host_management` where `host_management_id`=?", HostID).Scan(&host_management_id, &host_management_ip, &host_management_name, &host_management_user, &host_management_password, &host_management_key, &host_management_connect_method, &host_management_create_time, &host_management_modify_time)
	CheckError(err)
	host_management.HostManagementID = host_management_id
	host_management.HostManagementIP = host_management_ip
	host_management.HostManagementName = host_management_name
	host_management.HostManagementUser = host_management_user
	host_management.HostManagementPassword = host_management_password
	host_management.HostManagementKey = host_management_key
	host_management.HostManagementConnectMethod = host_management_connect_method
	host_management.HostManagementCreateTime = host_management_create_time
	host_management.HostManagementModifyTime = host_management_modify_time
	c.JSON(200, gin.H{
		"message":         "OK",
		"host_management": host_management,
	})
}

func DeleteHost(c *gin.Context) {
	HostID := c.PostForm("HostID")
	stmt, err := DB.Prepare("delete from `host_management` where `host_management_id`=?")
	CheckError(err)
	_, err = stmt.Exec(HostID)
	CheckError(err)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
