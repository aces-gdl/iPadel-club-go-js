package main

import (
	"fmt"
	"ipadel-club/initializers"
	"ipadel-club/server"
	"os"
	"time"
)

func init() {
	os.Setenv("TZ", "UTC")

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(" Error cargando timezone")
	}
	time.Local = loc
	// output current time zone
	fmt.Print("Local time zone ")
	fmt.Println(time.Now().Zone())
	fmt.Println(time.Now().Format("2006-01-02T15:04:05.000 CST"))

	initializers.LoadEnvVariables()
	initializers.ConnectTODBiPadelClub()
	if os.Getenv("DB_SYNC") == "YES" {
		initializers.SyncDatabase()
		fmt.Println("Sincronizando esquemas de base de datos...")
	}

}

func main() {
	server.Init()
}
