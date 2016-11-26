package config

import (
    "fmt"
    "flag"
    "github.com/BurntSushi/toml"
)

type tomlConfig struct {
    DB      database `toml:"database"`
    Server  server   `toml:"server"`
}

type database struct {
    Username string
    Password string
    Database string
}

type server struct {
    Port    int
}

func main() {

    var configPtr = flag.String("config","config.toml","path to configuration file")

    flag.Parse()

    var config tomlConfig
    if _, err := toml.DecodeFile(*configPtr, &config); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Database: %s %s %s\n",config.DB.Username,config.DB.Password,config.DB.Database)
    fmt.Printf("Port: %d\n",config.Server.Port)

}

