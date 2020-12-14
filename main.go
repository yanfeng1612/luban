package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"luban/service"
	"net/http"
	"strconv"
)

func main() {
	log.Println("start")
	//service.Workflow("/mnt/d/go-model/demo")
	addRouter()
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("启动服务异常", err)
	}
	log.Println("end")
}

func addRouter() {
	http.HandleFunc("/accCustInfo/getAccCustInfoById", func(writer http.ResponseWriter, request *http.Request) {
		id, _ := strconv.ParseInt(request.URL.Query().Get("id"), 10, 64)
		data, _ := json.Marshal(service.GetAccCustInfoById(id))
		writer.Write(data)
	})

	http.HandleFunc("/accCustInfo/add", func(writer http.ResponseWriter, request *http.Request) {
		bytes, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		acc := &service.AccCustInfo{}
		json.Unmarshal(bytes, &acc)
		log.Println(service.AddAccCustInfo(*acc))
		writer.Write([]byte("ok"))
	})

	http.HandleFunc("/project/list", func(writer http.ResponseWriter, request *http.Request) {
		pageNo, _ := strconv.Atoi(request.URL.Query().Get("pageNo"))
		pageSize, _ := strconv.Atoi(request.URL.Query().Get("pageSize"))
		res := service.GetProjectListBy(pageNo, pageSize)
		data, _ := json.Marshal(res)
		writer.Write(data)
	})

	http.HandleFunc("/project/init", func(writer http.ResponseWriter, request *http.Request) {
		id, _ := strconv.ParseInt(request.URL.Query().Get("id"), 10, 64)
		data, _ := json.Marshal(service.FirstInit(id))
		writer.Write(data)
	})

	http.HandleFunc("/project/assemble", func(writer http.ResponseWriter, request *http.Request) {
		id, _ := strconv.ParseInt(request.URL.Query().Get("id"), 10, 64)
		data, _ := json.Marshal(service.Assemble(id))
		writer.Write(data)
	})

	http.HandleFunc("/project/publish", func(writer http.ResponseWriter, request *http.Request) {
		id, _ := strconv.ParseInt(request.URL.Query().Get("id"), 10, 64)
		data, _ := json.Marshal(service.Publish(id))
		writer.Write(data)
	})

	http.HandleFunc("/project/start", func(writer http.ResponseWriter, request *http.Request) {
		id, _ := strconv.ParseInt(request.URL.Query().Get("id"), 10, 64)
		data, _ := json.Marshal(service.Start(id))
		writer.Write(data)
	})

	http.HandleFunc("/project/stop", func(writer http.ResponseWriter, request *http.Request) {
		id, _ := strconv.ParseInt(request.URL.Query().Get("id"), 10, 64)
		data, _ := json.Marshal(service.Stop(id))
		writer.Write(data)
	})

	http.HandleFunc("/project/restart", func(writer http.ResponseWriter, request *http.Request) {
		id, _ := strconv.ParseInt(request.URL.Query().Get("id"), 10, 64)
		data, _ := json.Marshal(service.Restart(id))
		writer.Write(data)
	})

}
