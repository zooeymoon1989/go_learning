package main

import (
	"net/http"
	"io"
)

func hello(w http.ResponseWriter , r *http.Request)  {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(w,
		`<DOCTYPE html>
			<html>
				<head>
					<title>Hello, World</title>
				</head>
				<body>
					Hello, World!
				</body>
			</html>`,
	)


}

func main() {
	http.HandleFunc("/hello",hello)
	http.ListenAndServe(":9090",nil)
}

//我们可以通过使用FileServer来处理静态文件
/*
http.Handle(
	"/assets/",
	http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
	),
)
 */
