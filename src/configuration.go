package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Server WebServer
	Client WebClient
	Memory Memory
	Disk   Disk
	Cpu    CPU
	Docker Docker
	Health HealthCheck
	Ping   PingCheck
}

func (s *Configuration) validate() {
	s.Server.validate()
	s.Client.validate()
	s.Cpu.validate()
	s.Disk.validate()
	s.Docker.validate()
	s.Memory.validate()
	s.Health.validate()
	s.Ping.validate()
}

var Conf Configuration

func init() {
	var config string
	if len(os.Args) == 2 {
		config = os.Args[1]
	} else {
		config = "./config.yml"
	}
	log.Printf("configuration file: %v\n", config)
	data, _ := ioutil.ReadFile(config)
	Conf = Configuration{}
	yaml.Unmarshal(data, &Conf)
	if !Conf.Memory.Enable && !Conf.Disk.Enable && !Conf.Cpu.Enable && !Conf.Docker.Enable && !Conf.Server.Enable {
		log.Fatal("have no monitor enabled,check your config.yml")
	}
	Conf.validate()
}

type WebClient struct {
	Enable bool
	Url    string
	Method string
}

func (s *WebClient) validate() {
	if s.Method == "" {
		s.Method = "POST"
	}
}

type WebServer struct {
	Enable bool
	Addr   string
	Token  string
}

func (s *WebServer) validate() {
	if len(s.Addr) == 0 {
		s.Addr = "0.0.0.0:80"
	}
}

type Memory struct {
	Enable   bool
	Limit    float64
	Frequcey uint
}

func (s *Memory) validate() {
	if s.Enable && s.Frequcey == 0 {
		s.Frequcey = 1000
	}
}

type Disk struct {
	Enable   bool
	All      bool
	Limit    float64
	Frequcey uint
	Paths    []DiskPath
}
type DiskPath struct {
	Path  string
	Limit float64
}

func (s *Disk) validate() {
	if s.Enable && s.Frequcey == 0 {
		s.Frequcey = 1000
	}
}

type CPU struct {
	Enable   bool
	Limit    float64
	Duration int
	Frequcey uint
}

func (s *CPU) validate() {
	if s.Enable && s.Frequcey == 0 {
		s.Frequcey = 1000
	}
	if s.Enable && s.Duration == 0 {
		s.Duration = 100
	}
}

type Docker struct {
	Enable     bool
	Frequcey   uint
	Containers []DockerContainer
}
type DockerContainer struct {
	Id   string
	Name string
}

func (s *Docker) validate() {
	if s.Enable && s.Frequcey == 0 {
		s.Frequcey = 1000
	}
}

type HealthCheck struct {
	Enable   bool
	Frequcey uint
	Urls     []string
}

func (s *HealthCheck) validate() {
	if s.Enable && s.Frequcey == 0 {
		s.Frequcey = 1000
	}
}

type PingCheck struct {
	Enable   bool
	Frequcey uint
	Count    uint
	Timeout  uint
	Ips      []string
}

func (s *PingCheck) validate() {
	if s.Enable && s.Frequcey == 0 {
		s.Frequcey = 1000
	}
	if s.Enable && s.Count == 0 {
		s.Count = 3
	}
	if s.Enable && s.Timeout == 0 {
		s.Timeout = 10
	}

}

/*func main() {
	text, _ := yaml.Marshal(Configuration{})
	log.Println("%v", string(text))
}
*/
