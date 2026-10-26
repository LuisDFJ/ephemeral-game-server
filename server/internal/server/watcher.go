package server

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const watcher_timeout time.Duration = 10 * time.Minute

func (s *Server) startWatcher() {
	kill := make(chan bool)
	done := make(chan bool)
	s.kill_watch = kill
	s.done_watch = done
	go s.watcher(kill, done)
}

func (s *Server) watcher(kill <-chan bool, done chan<- bool) {
	n := os.Getenv("VPS_EXPIRATION")
	exp, err := strconv.Atoi(n)
	if err != nil { exp = 3 }
	expiration := time.Duration(exp) * time.Hour
	fmt.Println("Expiration set to: ", expiration)
	loop: for {
		select {
		case <-kill:
			break loop
		case <-time.After(watcher_timeout):
			fmt.Println("Watcher timeout")
			vps, err := s.doc.GetDroplet()
			fmt.Println(vps)
			if err != nil || !vps.Alive { continue loop }
			timestamp, err := s.db.GetLastLog()
			if err != nil { continue loop }
			t := time.Unix(timestamp, 0).Add(expiration)
			if t.Before(time.Now()) {
				fmt.Println("VPS Expired ... Destroying")
				if err := s.doc.DelVps(vps.Id); err == nil {
					fmt.Println("Succesfully destroyed")
				}
			}
		}
	}
	done<- true
}
