/*
  Isana, a software for the game of Go
  Copyright (C) 2014 Tetsuo FUJII

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	. "github.com/foozea/isana/protocol"
)

// Defines engine name and version.
const (
	name    string = "Isana"
	version string = "0.1"
)

var (
	parallelNumber int
	trialNumber    int
)

func init() {
	// parse flags
	flag.IntVar(&parallelNumber, "parallels", runtime.NumCPU(), "parallel number")
	flag.IntVar(&parallelNumber, "p", runtime.NumCPU(), "parallel number")
	flag.IntVar(&trialNumber, "trials", 2000, "uct trial number")
	flag.IntVar(&trialNumber, "t", 2000, "uct trial number")
	flag.Parse()

	runtime.GOMAXPROCS(parallelNumber)

	Engine.Name = name
	Engine.Version = version
	Engine.Trials = trialNumber
}

func scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	// handle SIGNAL
	go func() {
		for sig := range c {
			println(sig)
			fmt.Printf("Interrupted...\n")
		}
	}()

	for {
		// parse input commands and handle them.
		input := strings.Split(scan(), " ")
		command := input[0]

		// first string is a command name.
		ArgsForHandlers = input[1:len(input)]
		if !Dispatcher.HasHandler(command) {
			fmt.Println("= unknown command")
			continue
		}
		Dispatcher.CallHandler(command)
	}
}
