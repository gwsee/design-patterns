package main

import "fmt"

/**
Nginx 这样的 Web 服务器可充当应用程序服务器的代理：

提供了对应用程序服务器的受控访问权限。
可限制速度。
可缓存请求。
*/
func main() {

	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	/**
	Url: /app/status
	HttpCode: 200
	Body: Ok

	Url: /app/status
	HttpCode: 200
	Body: Ok

	Url: /app/status
	HttpCode: 403
	Body: Not Allowed

	Url: /app/status
	HttpCode: 201
	Body: User Created

	Url: /app/status
	HttpCode: 404
	Body: Not Ok
	*/
}
