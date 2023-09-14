package flexdds

import (
	"fmt"
	"time"

	"github.com/rydyb/flexdds/ad9910"
	"github.com/rydyb/telnet"
)

type Config struct {
	Host     string
	Slot     uint16
	SysClock float64
}

// Client is a telnet client to the Client controller.
type Client struct {
	config Config
	client telnet.Client
}

// Open returns a new Client client for a controller with addr and DDS slot.
func Open(c Config) (*Client, error) {
	client := telnet.Client{
		Timeout: time.Duration(1) * time.Second,
		Address: fmt.Sprintf("%s:%d", c.Host, 26000+c.Slot),
	}

	err := client.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open tcp socket to %s: %w", client.Address, err)
	}

	out, err := client.Exec(fmt.Sprintf("75f4a4e10dd4b6b%d", c.Slot))
	if err != nil {
		return nil, fmt.Errorf("failed to send authentication token: %w", err)
	}
	if out != "Auth OK" {
		return nil, fmt.Errorf("failed to perform authentication: %s", out)
	}

	return &Client{
		config: c,
		client: client,
	}, nil
}

// Close closes the telnet connection.
func (c *Client) Close() error {
	return c.client.Close()
}

// Singletone configures channel to output a single frequency with relative amplitude.
func (c *Client) Singletone(ch uint8, ampl, freq float64) error {
	asf := ad9910.LogAmplScaleToASF(ampl)
	ftw := ad9910.FreqOutToFTW(freq, c.config.SysClock)

	if _, err := c.client.Exec(fmt.Sprintf("dcp %d spi:cfr2=0x01400820", ch)); err != nil {
		return fmt.Errorf("failed to configure CFR2 register: %w", err)
	}
	if _, err := c.client.Exec(fmt.Sprintf("dcp %d spi:stp0=0x%x0000%x", ch, asf, ftw)); err != nil {
		return fmt.Errorf("failed to configure STP0 register: %w", err)
	}
	if _, err := c.client.Exec(fmt.Sprintf("dcp %d update:u", ch)); err != nil {
		return fmt.Errorf("failed to update dds: %w", err)
	}

	return nil
}
