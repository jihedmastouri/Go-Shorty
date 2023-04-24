package service

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

const (
	ttl = time.Second * 8 // Time To Live
)

type service struct {
	agent      *api.Agent
	name       string
	id         string
	consulAddr string
}

type ServiceRegister int8

const (
	Consul ServiceRegister = iota
)

type ConfigProvider int8

const (
	ConsulConfig ConfigProvider = iota
)

type InitConfig struct {
	// Type Of Service Register
	// Default: Consul
	ServiceRegister ServiceRegister

	// Where to get the service Configs from
	// Default: ConsulConfig
	ConfigProvider ConfigProvider
}

// Initialize a new Service
// You must set Consul  Address as Environment variable `CONSUL_HTTP_ADDR`
func New(name string) *service {
	id := fmt.Sprintf("%s-%s", "my-service", uuid.NewString())
	consul := os.Getenv("CONSUL_HTTP_ADDR")

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal("New Service:", err)
	}

	agent := client.Agent()

	return &service{
		agent:      agent,
		id:         id,
		name:       name,
		consulAddr: consul,
	}
}

// For Future Compatibility
func (s *service) Init(c InitConfig) {}

// Start The Service:
// - Register The service with Consul
// - Listen on a tcp connection for Health updates
func (s *service) Start() {
	ownAddress, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
	}

	serviceDef := &api.AgentServiceRegistration{
		Name:    s.name,
		ID:      s.id,
		Address: ownAddress,
		Port:    50051,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: ttl.String(),
			TTL:                            ttl.String(),
			CheckID:                        s.id,
			TLSSkipVerify:                  true,
		},
	}

	if err := s.agent.ServiceRegister(serviceDef); err != nil {
		log.Fatal("Registeration:", err)
	}


	go s.keepAlive()
	go s.acceptLoop()

}

func (s *service) acceptLoop() {
	ln, err := net.Listen("tcp", ":8500")
	if err != nil {
		log.Fatal("Listening | Registeration:", err)
	}

	for {
		if _, err = ln.Accept(); err != nil {
			log.Fatal("Accept Loop | Registeration:", err)
		}
	}
}

// Update Health status before the Time To Live (TTL) expires
func (s *service) keepAlive() {
	ticker := time.NewTicker(ttl / 2)
	for {
		if err := s.agent.UpdateTTL(s.id, "online", api.HealthPassing); err != nil {
			log.Fatal(err)
		}
		<-ticker.C
	}
}
