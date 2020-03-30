package main

func main() {
	port := ":3000"

	server := NewServer(port)
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.addMiddleware(HandleHome, CheckAuth(), loggin()))
	server.Listen()
}
