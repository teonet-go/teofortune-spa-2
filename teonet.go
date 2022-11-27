package main

import (
	"net/http"

	"github.com/teonet-go/teonet"
)

const teofortune = "8agv3IrXQk7INHy5rVlbCxMWVmOOCoQgZBF"

// Teonet data and methods receiver
type Teonet struct {
	*teonet.Teonet
	fortune *teonet.APIClient
}

// Create new Teonet connection and connect to teofortune app
func newTeonet(appName string) (teo *Teonet, err error) {

	teo = new(Teonet)

	// Create teonet connection
	teo.Teonet, err = teonet.New(appName)
	if err != nil {
		return
	}
	// defer teo.Close()

	// Connect to teonet
	err = teo.Connect()
	if err != nil {
		return
	}

	// Connect to fortune server
	for i := 1; ; i++ {
		err = teo.ConnectTo(teofortune)
		switch {
		case err != nil && i >= 5:
			return
		case err != nil:
			continue
		}
		break
	}

	// Connect to fortune server api
	teo.fortune, err = teo.NewAPIClient(teofortune)
	if err != nil {
		return
	}

	return
}

// getFortune get fortune message from teonet
func (teo *Teonet) getFortune() (fortune string, err error) {

	// Get fortune message
	const fortaCmd = "forta"
	id, err := teo.fortune.SendTo(fortaCmd, nil)
	if err != nil {
		return
	}
	data, err := teo.fortune.WaitFrom(fortaCmd, uint32(id))
	if err != nil {
		return
	}
	fortune = string(data)

	return
}

// fortuneHandler get fortune message handler
func (teo *Teonet) fortuneHandler(w http.ResponseWriter, r *http.Request) {
	fortune, err := teo.getFortune()
	if err != nil {
		w.Write([]byte("can't get fortune, err: " + err.Error()))
	}
	w.Write([]byte(fortune))
}
