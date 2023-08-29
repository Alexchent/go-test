package demo

import "fmt"

type SLogMods struct {
	message chan []byte
}

var Slog = &SLogMods{}

func init() {
	Slog.message = make(chan []byte, 10000)
}

func (s *SLogMods) PreSend(msg []byte) {
	s.message <- msg
}

func (s *SLogMods) Send() {
	for {
		select {
		case msg := <-s.message:
			fmt.Println(string(msg))
		}
	}
}
