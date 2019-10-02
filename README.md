# linux-service

### Example

```go
package main

import (
	"time"

	service "github.com/bendersilver/linux-service"
)

type web struct{}

// StartService -
func (w *web) StartService() {}

// StopService -
func (w *web) StopService() {}

// Init -
func (w *web) Init() {}

func main() {
	service.Init("vt-web", "bot")
	service.Run(service.BotService(new(web)))
}


```