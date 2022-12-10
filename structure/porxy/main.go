package main

import "fmt"

// web服务应该具有处理请求的能力
type Server interface {
	handleRequest(url, method string) (int, string)
}

// web应用程序
type Application struct {
}

func (a Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 200, "User Created Success!"
	}
	return 404, "404 Not Found"
}

// nginx 代理web应用处理请求，做api接口请求限流
type NginxServer struct {
	application  Application
	MaxReqNum    int            // 最大请求数
	LimitRateMap map[string]int // 缓存每个接口的请求数
}

func NewNginxServer(app Application, max int) *NginxServer {
	return &NginxServer{
		application:  app,
		MaxReqNum:    max,
		LimitRateMap: make(map[string]int),
	}
}

// 代理web应用请求
func (n NginxServer) handleRequest(url, method string) (int, string) {
	if !n.checkReqRate(url) {
		return 403, "Not Allowed"
	}

	// 接口限流后转发请求到真实web应用
	return n.application.handleRequest(url, method)
}

// 接口限流和缓存
func (n *NginxServer) checkReqRate(url string) bool {
	reqNum := n.LimitRateMap[url]

	if reqNum >= n.MaxReqNum {
		return false
	}
	n.LimitRateMap[url]++

	return true
}

func main() {

	nginx := NewNginxServer(Application{}, 2)
	respCode, respBody := nginx.handleRequest("/app/status", "GET")
	fmt.Printf("URL:%s \n返回状态码:%d,响应内容:%s \n\n", "/app/status", respCode, respBody)

	respCode, respBody = nginx.handleRequest("/app/status", "GET")
	fmt.Printf("URL:%s \n返回状态码:%d,响应内容:%s \n\n", "/app/status", respCode, respBody)

	// 超过了最大限流数 返回403
	respCode, respBody = nginx.handleRequest("/app/status", "GET")
	fmt.Printf("URL:%s \n返回状态码:%d,响应内容:%s \n\n", "/app/status", respCode, respBody)

	respCode, respBody = nginx.handleRequest("/create/user", "POST")
	fmt.Printf("URL:%s \n返回状态码:%d,响应内容:%s \n\n", "/create/user", respCode, respBody)

}

/* output
URL:/app/status
返回状态码:200,响应内容:Ok

URL:/app/status
返回状态码:200,响应内容:Ok

URL:/app/status
返回状态码:403,响应内容:Not Allowed

URL:/create/user
返回状态码:200,响应内容:User Created Success!

*/
