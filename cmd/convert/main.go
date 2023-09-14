package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/rydyb/flexdds/ad9910"
)

type Ctx struct{}

type FreqOutCmd struct {
	SysClock float64 `name:"system-clock" default:"1e9" help:"The system' clocks frequency in Hz."`
	FreqOut  float64 `arg:"" name:"frequency" help:"The output frequency in Hz."`
}

func (cmd *FreqOutCmd) Run(ctx *Ctx) error {
	fmt.Printf("0x%x\n", ad9910.FreqOutToFTW(cmd.FreqOut, cmd.SysClock))
	return nil
}

type LinAmplScaleCmd struct {
	AmplScale float64 `arg:"" name:"amplitude" help:"The linear amplitude scale from 0.0 to 1.0."`
}

func (cmd *LinAmplScaleCmd) Run(ctx *Ctx) error {
	fmt.Printf("0x%x\n", ad9910.LinAmplScaleToASF(cmd.AmplScale))
	return nil
}

type LogAmplScaleCmd struct {
	AmplScale float64 `arg:"" name:"amplitude" help:"The amplitude scale in dBm relative to maximum output."`
}

func (cmd *LogAmplScaleCmd) Run(ctx *Ctx) error {
	fmt.Printf("0x%x\n", ad9910.LogAmplScaleToASF(cmd.AmplScale))
	return nil
}

var cli struct {
	FreqOut      FreqOutCmd      `cmd:"freq-out" help:"Convert an output frequency in Hz to FTW register value."`
	LogAmplScale LogAmplScaleCmd `cmd:"log-ampl-scale" help:"Convert logarithmic amplitude scale in dBm to ASF register value."`
	LinAmplScale LinAmplScaleCmd `cmd:"lin-ampl-scale" help:"Convert linear amplitude scale from 0.0 to 1.0 to ASF register value."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Ctx{})
	ctx.FatalIfErrorf(err)
}
