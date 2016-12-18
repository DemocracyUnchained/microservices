package main

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
    Host string
    Port int
}

type server struct {
    Port    int
}

var Config tomlConfig

func loadConfig() {

    var configPtr = flag.String("config","/etc/config.toml","path to configuration file")

    flag.Parse()

    if _, err := toml.DecodeFile(*configPtr, &Config); err != nil {
        fmt.Println(err)
        return
    }

}

