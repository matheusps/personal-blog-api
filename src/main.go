package main

func main() {
	server := provideServer()
	AutoMigrateTypes()
	server.Run(":8080")
}
