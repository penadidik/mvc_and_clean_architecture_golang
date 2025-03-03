package main

import (
    "backend_architecture_golang/database"
    "backend_architecture_golang/routes"
)

func main() {
    database.InitDB()
    r := routes.SetupRouter()
    r.Run(":8080")
}
