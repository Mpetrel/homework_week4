package main

import (
	"context"
	"flag"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v2"
	"homework_week4/internal/conf"
	"homework_week4/internal/data"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var flagConf string
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.Parse()

	log.Println(flagConf)
	yamlFile, err := ioutil.ReadFile(flagConf)
	if err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err = yaml.Unmarshal(yamlFile, &bc); err != nil {
		panic(err)
	}
	// setup db
	err = data.SetupDB(bc.GetData())
	if err != nil {
		panic(err)
	}

	// create service
	lis, err := net.Listen("tcp", bc.GetServer().GetGrpc().GetAddr())
	if err != nil {
		panic(err)
	}
	s, err := initServer(bc.GetServer(), bc.GetData())
	if err != nil {
		panic(err)
	}

	// run service by errgroup
	group, ctx := errgroup.WithContext(context.Background())
	// listen system event
	group.Go(func() error {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <- ctx.Done():
			return ctx.Err()
		case sig := <- c:
			return errors.Errorf("receive system signal: %v", sig)
		}
	})

	group.Go(func() error {
		select {
		case <-ctx.Done():
			log.Printf("errgroup finished, close server..")
		}
		// stop server
		s.Stop()
		return nil
	})

	group.Go(func() error {
		log.Printf("start rpc server at: %s", bc.GetServer().GetGrpc().GetAddr())
		err := s.Serve(lis)
		return err
	})

	if err = group.Wait(); err != nil {
		panic(err)
	}

}
