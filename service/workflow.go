package service

import (
	"log"
	"luban/utils"
	"os"
	"strings"
	"time"
)

const (
	GIT_ROOT_PATH             = "d://luban-repository//"
	PACKAGE_ROOT_PATH         = "d://luban-repository//package//"
	PACKAGE_MACHINE_ROOT_PATH = "d://export//package//"
	TIME_VERSION_FORMAT       = "20060102-150405"
)

type WorkflowDomain struct {
	rootPath string
}

func Assemble(id int64) Response {
	project := GetProjectById(id)
	err := os.Chdir(GIT_ROOT_PATH + project.Name)
	if err != nil {
		return NewFailResponse()
	}

	arr := strings.Split(project.AssembleCommand, " ")
	args := []string{}
	for _, s := range arr {
		if s != "" {
			args = append(args, s)
		}
	}

	utils.ExecuteCmd("git", "pull", "origin", "master")
	utils.ExecuteCmd("mvn", args...)
	version := time.Now().Format(TIME_VERSION_FORMAT)

	destPath := PACKAGE_ROOT_PATH + project.Name + "/" + version
	utils.ExecuteCmd("mkdir", destPath)
	utils.ExecuteCmd("cp", "-r", GIT_ROOT_PATH+project.Name+"/"+project.AssemblePath+"/*", destPath+"/")
	AddPublishHistory(project.Id, version)
	return NewSuccessResponse()
}

func Publish(id int64) Response {
	project := GetProjectById(id)
	version := GetLastVersion(id)
	sourcePath := PACKAGE_ROOT_PATH + project.Name + "/" + version + "/"
	destPath := PACKAGE_MACHINE_ROOT_PATH + project.Name + "/" + version + "/"
	ipList := GetMachineIpListBy(id)
	for _, ip := range ipList {
		if ip == "127.0.0.1" || ip == "localhost" {
			utils.ExecuteCmd("mkdir", "-p", destPath)
			utils.ExecuteCmd("cp", "-r", sourcePath+"*", destPath)
		}
	}

	//go Restart(id)
	return NewSuccessResponse()
}

func Start(id int64) Response {
	project := GetProjectById(id)
	err := os.Chdir(PACKAGE_ROOT_PATH + project.Name + "/" + GetLastVersion(id) + "/bin")
	if err != nil {
		return NewFailResponse()
	}
	go utils.ExecuteCmd("sh", "start.sh")
	return NewSuccessResponse()
}

func Stop(id int64) Response {
	project := GetProjectById(id)
	err := os.Chdir(PACKAGE_ROOT_PATH + project.Name + "/" + GetLastVersion(id) + "/bin")
	if err != nil {
		return NewFailResponse()
	}
	go utils.ExecuteCmd("sh", "stop.sh")
	return NewSuccessResponse()
}

func Restart(id int64) Response {
	Stop(id)
	Start(id)
	return NewSuccessResponse()
}

func FirstInit(id int64) Response {
	project := GetProjectById(id)
	err := os.Chdir(GIT_ROOT_PATH)
	if err != nil {
		log.Fatal("路径错误 gitRootPath: " + GIT_ROOT_PATH)
		return NewFailResponse()
	}
	utils.ExecuteCmd("git", "clone", project.GitUrl)
	return NewSuccessResponse()
}
