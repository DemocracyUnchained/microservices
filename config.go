import   "github.com/BurntSushi/toml"

var tomlBlob = `
# Some comments.
[alpha]
ip = "10.0.0.1"

	[alpha.config]
	Ports = [ 8001, 8002 ]
	Location = "Toronto"
	Created = 1987-07-05T05:45:00Z

[beta]
ip = "10.0.0.2"

	[beta.config]
	Ports = [ 9001, 9002 ]
	Location = "New Jersey"
	Created = 1887-01-05T05:55:00Z
`

type serverConfig struct {
    Ports    []int
    Location string
    Created  time.Time
}

type server struct {
    IP     string       `toml:"ip,omitempty"`
    Config serverConfig `toml:"config"`
}

type servers map[string]server

var config servers
if _, err := Decode(tomlBlob, &config); err != nil {
    log.Fatal(err)
}

for _, name := range []string{"alpha", "beta"} {
    s := config[name]
    fmt.Printf("Server: %s (ip: %s) in %s created on %s\n",
        name, s.IP, s.Config.Location,
        s.Config.Created.Format("2006-01-02"))
    fmt.Printf("Ports: %v\n", s.Config.Ports)
}

