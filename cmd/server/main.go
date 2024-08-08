package main

import (
	"access-point/app"
)

func main() {
	app := app.NewApplication() // Create a new application instance
	app.Init()                  // Initialize the application (e.g., set up configurations, dependencies)
	app.Run()                   // Run the application (e.g., start the server, process requests)
	app.Wait()                  // Optionally, wait for the application to finish (e.g., handle signals)
	app.Cleanup()               // Clean up resources (e.g., close connections, clean up temporary files)
}