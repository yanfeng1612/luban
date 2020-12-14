package service

import (
	"log"
	"luban/utils"
)

type AccCustInfo struct {
	Id     int64  `json:"id"`
	Uuid   string `json:"uuid"`
	CustNo string `json:"custNo"`
}

func GetAccCustInfoById(id int64) *AccCustInfo {
	list := utils.QuerySql("SELECT * FROM acc_cust_info WHERE id = ?", id)
	if len(list) != 1 {
		return nil
	}
	m := list[0]
	return &AccCustInfo{m["id"].(int64), m["uuid"].(string), m["cust_no"].(string)}
}

func AddAccCustInfo(info AccCustInfo) bool{
	log.Println(info.Id)
	utils.ExecuteSql("INSERT INTO acc_cust_info (uuid,cust_no,pin,prof_flag,version,yn,created_time,modified_time,biz_date) VALUES (?,?,'123',1,1,1,NOW(),NOW(),'2020-11-24')",info.Uuid,info.CustNo)
	return true
}
