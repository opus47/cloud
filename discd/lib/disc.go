package discd

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type Track struct {
	Length int // in seconds
}

type DiscInfo struct {
	Tracks []Track
}

func (d DiscInfo) String() string {
	s := ""
	s += fmt.Sprintf("%d tracks\n", len(d.Tracks))
	s += strings.Repeat("-", len(s)-1) + "\n"
	for i, t := range d.Tracks {
		s += fmt.Sprintf("Track%d - %d\n", i, t.Length)
	}
	return s
}

func ReadDiscInfo() (*DiscInfo, error) {

	cmd := exec.Command("cdparanoia", "-qJQ")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("[read-info] exec error: %s", string(out))
	}

	info := &DiscInfo{}

	err = json.Unmarshal(out, info)
	if err != nil {
		return nil, fmt.Errorf("[read-info] error parsing ouput: %s", err)
	}

	return info, nil

}

func RipTrack(i int) {

	// set up progress socket
	sockfile := fmt.Sprintf("/tmp/o47p-%d.sock", os.Getpid())
	ln, err := net.ListenUnixgram(
		"unixgram",
		&net.UnixAddr{sockfile, "unixgram"},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening @ %s", sockfile)

	// handle system signals - shutdown progress socket on sig
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

	go func(ln *net.UnixConn, c chan os.Signal) {
		sig := <-c
		log.Printf("caught signal %s: shutting down.", sig)
		ln.Close()
		os.Exit(0)
	}(ln, sigc)

	// report progress in separate thread
	go func(ln *net.UnixConn, i int) {
		for {
			buf := make([]byte, 4)
			_, err = ln.Read(buf)
			if err != nil {
				log.Fatal(err)
				return
			}

			bits := binary.LittleEndian.Uint32(buf)
			value := math.Float32frombits(bits)

			fmt.Printf("track %d = %f\r", i, value)
		}
	}(ln, i)

	// call ripping command and wait for finish
	cmd := exec.Command(
		"/usr/local/bin/cdparanoia",
		"-u", sockfile,
		"-B",
		"--",
		strconv.FormatInt(int64(i), 10),
	)
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	ln.Close()

}
