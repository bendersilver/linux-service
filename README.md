# linux-service

### Example

```go
package main

import (
	"time"

	service "github.com/bendersilver/linux-service"
)

type Example struct {
    exit chan struct{}
}

func (s *Example) Start() {
	s.exit = make(chan struct{})
	ticker := time.Tick(time.Second)
	for {
		select {
		case <-ticker:
			logger.Info("ticker")
		case <-s.exit:
			break
		}
	}
}

func (s *Example) Stop() {
	close(s.exit)
}

func main() {
	service.SetName("my-tiker-example")
	service.Run(service.Ifce(new(serv)))
}

```