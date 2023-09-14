package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/rydyb/flexdds"
)

var cli struct {
	Host       string  `name:"host" required:"" help:"The hostname or ip address of the FlexDDS controller."`
	Slot       uint16  `name:"slot" required:"" help:"The number of the FlexDDS slot from 0 to 5."`
	Channel    uint8   `name:"channel" required:"" help:"The number of the FlexDDS channel 0 or 1."`
	SysClock   float64 `name:"system-clock" default:"1e9" help:"The system' clocks frequency in Hz."`
	Singletone struct {
		Amplitude float64 `name:"amplitude" help:"The singletone amplitude in dBm."`
		Frequency float64 `name:"frequency" required:"" help:"The frequency of the singletone in Hz."`
	} `cmd:"singletone" help:"Configure a singletone output."`
}

func main() {
	args := kong.Parse(&cli)

	if cli.Slot > 5 {
		log.Fatalf("Slot number cannot be greater than five.")
	}
	if cli.Channel > 1 {
		log.Fatalf("Channel number has to be zero or one.")
	}

	flexdds, err := flexdds.Open(flexdds.Config{
		Host:     cli.Host,
		Slot:     cli.Slot,
		SysClock: cli.SysClock,
	})
	if err != nil {
		log.Fatalf("failed to open connection to %s: %s", cli.Host, err)
	}
	defer flexdds.Close()

	switch args.Command() {
	case "singletone":
		if err := flexdds.Singletone(cli.Channel, cli.Singletone.Amplitude, cli.Singletone.Frequency); err != nil {
			log.Fatalf("failed to configure channel %d to singletone with amplitude %f and frequency %f: %s", cli.Channel, cli.Singletone.Amplitude, cli.Singletone.Frequency, err)
		}
	}
}
