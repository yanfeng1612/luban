package service

import (
	"luban/utils"
)

type Project struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	ChineseName     string `json:"chineseName"`
	Principal       string `json:"principal"`
	Url             string `json:"url"`
	Type            int    `json:"type"`
	DevLang         int    `json:"devLang"`
	GitUrl          string `json:"gitUrl"`
	AssemblePath    string `json:"assemblePath"`
	AssembleCommand string `json:"assembleCommand"`
}

func GetProjectById(id int64) *Project {
	m := utils.QuerySingle("SELECT * FROM project WHERE id = ?", id)
	return &Project{m["id"].(int64), m["name"].(string), m["chinese_name"].(string), m["principal"].(string), m["url"].(string), m["type"].(int), m["dev_lang"].(int), m["git_url"].(string), m["assemble_path"].(string), m["assemble_command"].(string)}
}

func GetProjectListBy(pageNo int, pageSize int) Response {
	list := utils.QuerySql("SELECT * FROM project ORDER BY id DESC LIMIT ?,?", (pageNo-1)*pageSize, pageSize)
	data := make([]Project, 0)
	for _, m := range list {
		project := &Project{m["id"].(int64), m["name"].(string), m["chinese_name"].(string), m["principal"].(string), m["url"].(string), m["type"].(int), m["dev_lang"].(int), m["git_url"].(string), m["assemble_path"].(string), m["assemble_command"].(string)}
		data = append(data, *project)
	}

	count := utils.QuerySingle("SELECT COUNT(*) AS count FROM project")["count"]

	return NewSuccessResponseWithData(NewPageResult(pageNo, pageSize, count.(int64), data))
}

func AddPublishHistory(projectId int64, version string) {
	utils.ExecuteSql("INSERT INTO project_publish_history (project_id,publish_version,create_time) VALUES (?,?,NOW())", projectId, version)
}

func GetLastVersion(projectId int64) string {
	m := utils.QuerySingle("SELECT publish_version AS version FROM project_publish_history WHERE project_id = ? ORDER BY id DESC LIMIT 1",projectId)
	if len(m) == 1 {
		return m["version"].(string)
	}
	return ""
}

func GetMachineIpListBy(id int64) (ipList []string) {
	ipListMap := utils.QuerySql("SELECT ip FROM machine WHERE yn = 1 AND project_id = ?",id)
	for _, ipMap := range ipListMap {
		ipList = append(ipList,ipMap["ip"].(string))
	}
	return ipList
}
