package main

import (
	////"bufio"
	//"fmt"
	////"github.com/hpcloud/tail"
	"log"
	////"os"
	//"io"
	"net/http"
	//"time"
	_ "net/http/pprof"
)

//import "fmt"

//func HelloServer(w http.ResponseWriter, req *http.Request) {
//	io.WriteString(w, "hello, world!\n")
//}
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	//log.Println("start")
	//
	//go func() {
	//	http.HandleFunc("/hello", HelloServer)
	//	log.Println("zhixingalalll")
	//	err := http.ListenAndServe("127.0.0.1:8090", nil)
	//
	//	//log.Fatal(http.ListenAndServe("127.0.0.1:8090", nil))
	//	log.Fatal(err)
	//	log.Println("zhixingalalll..............")
	//}()
	//
	//log.Println("main runing .......")
	exit := make(chan bool)
	<-exit
	//fmt.Println(time.Now().String())
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//t, _ := tail.TailFile("/Users/zhangyachuan/Documents/proxy_access.log", tail.Config{Poll: true, Follow: true})
	//for line := range t.Lines {
	//	log.Println(line.Text)
	//	//fmt.Println(line.Text)
	//}
	//ReadLogfile("/Users/zhangyachuan/Documents/proxy_access.log")
}
//$remote_addr|$remote_user|[$time_local]|$host|$request|$status|$body_bytes_sent|$http_referer|'
//                     '$http_user_agent|$http_x_forwarded_for|$request_time|$upstream_response_time|'
//                     '$upstream_connect_time|$upstream_header_time|$upstream_http_via|$upstream_addr|'
//                     '$upstream_http_x_e_reqid|$upstream_http_x_m_reqid'

//103.20.32.163|-|[03/May/2017:17:33:45 +0800]|s1.cdn.xiangha.com|GET /favicon.ico HTTP/2.0|401|399|https://s1.cdn.xiangha.com/caipu/201507/0417/041728029421.jpg/MjgweDIyMA|Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36|-|0.196|0.196|0.000|0.196|http/1.1 cdn-cnc-gddg-dg-8 (ApacheTrafficServer/6.2.0 [cMsSf ])|127.0.0.1:8080|1493804025673921299:4639915|5lgAAEs4eMDYDrsU


//
//func ReadLogfile(logpath string) {
//	var debugLine = ""
//	// 读取配置文件
//	logfile, fileerr := os.Open(logpath)
//	if fileerr != nil {
//		log.Fatalln(fileerr)
//	}
//	defer logfile.Close()
//
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println(string(debugLine), err)
//
//			panic(err)
//		}
//	}()
//	rblog := bufio.NewReader(logfile)
//	for {
//		time.Sleep(1)
//		readline, _, nerr := rblog.ReadLine()
//		if nerr != nil {
//			log.Fatalln("----------------------" + nerr.Error())
//		}
//		debugLine = string(readline)
//		log.Println(debugLine)
//	}
//
//}
