package main

const port string = ":8080"

func main() {
	server := provideServer()
	AutoMigrateTypes()
	server.Run(port)
}
