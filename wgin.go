package wgin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func GinDefault(router *gin.Engine) {
	router.GET("/api/confg/AddControl", AddControl)
	router.GET("/api/confg/DelControl", DelControl)
	router.GET("/api/confg/GetControl", GetControl)
}

//添加控制
func AddControl(c *gin.Context) {
	suid, ok := c.GetQuery("uid")
	if !ok {
		return
	}
	types, _ := c.GetQuery("type")
	name := "control.json"
	if types == "1" {
		name = "wcontrol.json"
	}
	uid, err := strconv.Atoi(suid)
	if err != nil {
		return
	}
	var uids []int
	data, err := ioutil.ReadFile("./config/" + name)
	if err != nil {
		//return
	}
	err = json.Unmarshal(data, &uids)
	if err != nil {
		//return
	}
	uids = append(uids, uid)
	data, err = json.Marshal(uids)
	if err != nil {
		return
	}
	file, err := os.OpenFile("./config/"+name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	file.Write(data)
	file.Close()
	c.String(http.StatusOK, string(data))
}

//删除控制
func DelControl(c *gin.Context) {
	suid, ok := c.GetQuery("uid")
	if !ok {
		return
	}
	types, _ := c.GetQuery("type")
	name := "control.json"
	if types == "1" {
		name = "wcontrol.json"
	}
	uid, err := strconv.Atoi(suid)
	if err != nil {
		return
	}
	var uids []int
	data, err := ioutil.ReadFile("./config/" + name)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &uids)
	if err != nil {
		return
	}
	//删除用户
topDel:
	for i := range uids {
		if uids[i] == uid {
			uids = append(uids[:i], uids[i+1:]...)
			goto topDel
		}
	}
	data, err = json.Marshal(uids)
	if err != nil {
		return
	}
	file, err := os.OpenFile("./config/"+name, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	file.Write(data)
	file.Close()
	c.String(http.StatusOK, string(data))
}

//查询全部控制
func GetControl(c *gin.Context) {
	types, _ := c.GetQuery("type")
	name := "control.json"
	if types == "1" {
		name = "wcontrol.json"
	}
	data, err := ioutil.ReadFile("./config/" + name)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(data))
}
