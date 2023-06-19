package main

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Server struct {
	state State
	// config is the configuration of the game.
	config Config
	// conn is a hashset of existing connections.
	conn map[*Conn]struct{}
	// update is set to true when there are updates to be pushed to the client.
	update bool

	mu *sync.Mutex
}

type Config struct {
	SideA  string  `json:"side_a"`
	SideB  string  `json:"side_b"`
	Phases []Phase `json:"phases"`
}

type Phase struct {
	Text string `json:"text"`
	Time []int  `json:"time"`
}

type State struct {
	Title  string `json:"title"`
	Phase  string `json:"phase"`
	SideA  string `json:"side_a"`
	SideB  string `json:"side_b"`
	Timer  []int  `json:"timer"`
	Active int    `json:"active"`
	Paused bool   `json:"paused"`
}

// NewServer returns a server with the specified config.
func NewServer(cfg Config) *Server {
	return &Server{
		state: State{
			Timer: []int{},
		},
		config: cfg,
		conn:   make(map[*Conn]struct{}, 2),
		update: false,
		mu:     new(sync.Mutex),
	}
}

// Server registers the connection to the game.
func (s *Server) Register(c *Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.conn[c] = struct{}{}
}

// Loop is the game loop.
func (s *Server) Loop() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		s.tick()
	}
}

func (s *Server) tick() {
	s.mu.Lock()
	defer s.mu.Unlock()
	toClose := make([]*Conn, 0)
	for conn := range s.conn {
		if conn.closed {
			toClose = append(toClose, conn)
		}
	}
	for _, conn := range toClose {
		delete(s.conn, conn)
	}
	if !s.state.Paused && s.state.Active >= 0 && s.state.Active < len(s.state.Timer) {
		if s.state.Timer[s.state.Active] > 0 {
			s.state.Timer[s.state.Active]--
			s.update = true
		}
	}
	if !s.update {
		return
	}
	s.update = false
	for conn := range s.conn {
		s.pushState(conn)
	}
}

// ProcessCommand processes the specified commands.
func (s *Server) ProcessCommand(commands []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(commands) == 0 {
		return
	}
	switch commands[0] {
	case "title":
		s.update = true
		s.state.Title = strings.Join(commands[1:], " ")
	case "sideA":
		s.update = true
		s.state.SideA = strings.Join(commands[1:], " ")
	case "sideB":
		s.update = true
		s.state.SideB = strings.Join(commands[1:], " ")
	case "go":
		if len(commands) < 2 {
			log.Println("Expected arguments:", strings.Join(commands, " "))
			return
		}
		nums := make([]int, len(commands)-1)
		for i := 1; i < len(commands); i++ {
			var err error
			nums[i-1], err = strconv.Atoi(commands[i])
			if err != nil || nums[i-1] < 0 {
				log.Println("Unknown argument:", strings.Join(commands, " "))
				return
			}
		}
		if nums[0] >= len(s.config.Phases) {
			log.Println("Invalid argument 1:", strings.Join(commands, " "))
			return
		}
		s.update = true
		phase := s.config.Phases[nums[0]]
		s.state.Phase = phase.Text
		s.state.Timer = nums[1:]
		s.state.Paused = true
		s.state.Active = 0
	case "switchSide":
		max := len(s.state.Timer)
		if max <= 1 {
			log.Println("Cannot switch side when timer is of length", max)
			return
		}
		s.update = true
		s.state.Active++
		s.state.Active %= max
	case "togglePause":
		s.update = true
		s.state.Paused = !s.state.Paused
	default:
		log.Println("Unknown command:", strings.Join(commands, " "))
	}
}
