package flexdds

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/rydyb/flexdds/ad9910"
	"github.com/rydyb/telnet"
)

// FlexDDS is a telnet client to the FlexDDS controller.
type FlexDDS struct {
	SysClock float64
	client   telnet.Client
}

// Open returns a new FlexDDS client for a controller with addr and DDS slot.
func Open(host string, slot uint16) (*FlexDDS, error) {
	client := telnet.Client{
		Timeout: time.Duration(1) * time.Second,
		Address: fmt.Sprintf("%s:%d", host, 26000+slot),
	}

	err := client.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open tcp socket to %s: %w", client.Address, err)
	}

	out, err := client.Exec(fmt.Sprintf("75f4a4e10dd4b6b%d", slot))
	if err != nil {
		return nil, fmt.Errorf("failed to send authentication token: %w", err)
	}
	if out != "Auth OK" {
		return nil, fmt.Errorf("failed to perform authentication: %s", out)
	}

	return &FlexDDS{
		SysClock: 1e9,
		client:   client,
	}, nil
}

// Close closes the telnet connection.
func (c *FlexDDS) Close() error {
	return c.client.Close()
}

// Singletone configures channel to output a single frequency with relative amplitude.
func (c *FlexDDS) Singletone(ch uint8, ampl, freq float64) error {
	asf := ad9910.LogAmplScaleToASF(ampl)
	log.Debug().Msgf("asf: %X", asf)
	ftw := ad9910.FreqOutToFTW(freq, c.SysClock)
	log.Debug().Msgf("ftw: %x", ftw)

	if _, err := c.client.Exec(fmt.Sprintf("dcp %d spi:cfr2=0x01400820", ch)); err != nil {
		return fmt.Errorf("failed to configure CFR2 register: %w", err)
	}
	if _, err := c.client.Exec(fmt.Sprintf("dcp %d spi:stp0=%x0000%x", ch, asf, ftw)); err != nil {
		return fmt.Errorf("failed to configure STP0 register: %w", err)
	}
	if _, err := c.client.Exec(fmt.Sprintf("dcp %d update:u", ch)); err != nil {
		return fmt.Errorf("failed to update dds: %w", err)
	}

	return nil
}
