package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/meshplus/premo/internal/repo"

	"github.com/gobuffalo/packr"
	"github.com/meshplus/premo/internal/bitxhub"
	"github.com/urfave/cli/v2"
)

var testCMD = &cli.Command{
	Name:  "test",
	Usage: "test bitxhub function",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "concurrent",
			Aliases: []string{"c"},
			Value:   100,
			Usage:   "concurrent number",
		},
		&cli.IntFlag{
			Name:    "tps",
			Aliases: []string{"t"},
			Value:   500,
			Usage:   "all tx number",
		},
		&cli.IntFlag{
			Name:    "duration",
			Aliases: []string{"d"},
			Value:   60,
			Usage:   "test duration",
		},
		&cli.StringFlag{
			Name:    "key_path",
			Aliases: []string{"k"},
			Usage:   "Specify key path",
		},
		&cli.StringSliceFlag{
			Name:    "remote_bitxhub_addr",
			Aliases: []string{"r"},
			Usage:   "Specify remote bitxhub address",
			Value:   cli.NewStringSlice("localhost:60011"),
		},
		&cli.StringFlag{
			Name:  "type",
			Usage: "Specify tx type: interchain, data, transfer",
			Value: "transfer",
		},
		&cli.StringFlag{
			Name:  "appchain",
			Usage: "Specify appchain type: fabric, flato",
			Value: "flato",
		},
	},
	Action: benchmark,
}

func benchmark(ctx *cli.Context) error {
	box := packr.NewBox(repo.ConfigPath)
	//val, err := box.Find("fabric.validators")
	val, err := box.Find("single_validator")
	if err != nil {
		return err
	}
	proof, err := box.Find("proof_1.0.0_rc")
	if err != nil {
		return err
	}
	typ := ""

	appchain := ctx.String("appchain")
	if appchain == "fabric" {
		//val, err = box.Find("validator_fabric")
		//if err != nil {
		//	return err
		//}
		//proof, err = box.Find("proof_fabric")
		//if err != nil {
		//	return err
		//}
		typ = "Fabric V1.4.3"
	} else if appchain == "flato" {
		//val, err = box.Find("validator_flato")
		//if err != nil {
		//	return err
		//}
		//proof, err = box.Find("proof_flato")
		//if err != nil {
		//	return err
		//}
		typ = "Flato V1.0.3"
	} else {
		return fmt.Errorf("unsupported appchain type")
	}
	contract, err := box.Find("rule.wasm")
	if err != nil {
		return err
	}

	keyPath := ctx.String("key_path")
	if keyPath == "" {
		//default use node4
		keyPath, err = repo.Node4Path()
		if err != nil {
			return err
		}
	}
	config := &bitxhub.Config{
		Concurrent:  ctx.Int("concurrent"),
		TPS:         ctx.Int("tps"),
		Duration:    ctx.Int("duration"),
		Type:        ctx.String("type"),
		KeyPath:     keyPath,
		BitxhubAddr: ctx.StringSlice("remote_bitxhub_addr"),
		Validator:   string(val),
		Proof:       proof,
		Rule:        contract,
		Appchain:    typ,
	}

	if config.Concurrent > config.TPS {
		return fmt.Errorf("error: concurrent should be less than tps")
	}

	broker, err := bitxhub.New(config)
	if err != nil {
		return err
	}

	handleShutdown(broker)

	err = broker.Start(config.Type)
	if err != nil {
		return err
	}

	return nil
}

func handleShutdown(node *bitxhub.Broker) {
	current := time.Now()
	var stop = make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)
	go func() {
		<-stop
		fmt.Println("received interrupt signal, shutting down...")
		if err := node.Stop(current); err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
}
