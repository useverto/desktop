package main

import (
	"log"
	"os"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/kardianos/osext"
)

func setupWatcher() (chan struct{}, error) {
	file, err := osext.Executable()
	if err != nil {
		return nil, err
	}
	log.Printf("watching %q\n", file)
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	done := make(chan struct{})
	go func() {
		for {
			select {
			case e := <-w.Events:
				log.Printf("watcher received: %+v", e)
				err := syscall.Exec(file, os.Args, os.Environ())
				if err != nil {
					log.Fatal(err)
				}
			case err := <-w.Errors:
				log.Printf("watcher error: %+v", err)
			case <-done:
				log.Print("watcher shutting down")
				return
			}
		}
	}()
	err = w.Add(file)
	if err != nil {
		return nil, err
	}
	return done, nil
}
