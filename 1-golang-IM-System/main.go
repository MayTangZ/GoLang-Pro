package main

func main() {

	// Use the NewServer function to create a new Server instance
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
