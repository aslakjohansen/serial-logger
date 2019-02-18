package main

import (
    "fmt"
    "os"
    "time"
    "github.com/tarm/serial"
)

const BUFSIZE = 1024

var buf []byte = make([]byte, BUFSIZE)

func append2log (log *os.File, t0 time.Time, t1 time.Time, line string) {
    var fullline string = fmt.Sprintf("%d.%09d,%d.%09d,%s\n", t0.Unix(), t0.Nanosecond(), t1.Unix(), t1.Nanosecond(), line)
    
    if _, err := log.Write([]byte(fullline)) ; err != nil {
        fmt.Println(err)
        os.Exit(3)
    }
}

func main () {
    // guard: command line args
    if len(os.Args) != 3 {
        fmt.Printf("Syntax: %s DEVICE FILENAME\n", os.Args[0])
        fmt.Printf("        %s /dev/ttyACM0 log.csv\n", os.Args[0])
        os.Exit(1)
    }
    var dev_path string = os.Args[1]
    var log_path string = os.Args[2]
    
    // print out configuration
    fmt.Printf("serial-logger: %s -> %s\n", dev_path, log_path)
    
    // open log file
    log, err := os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    
    // open serial port
    c := &serial.Config{Name: dev_path, Baud: 9600}
    s, err := serial.OpenPort(c)
    if err != nil {
        fmt.Println(err)
    }
    
    s.Read(buf)
    
    // dummy write to appease golang
    append2log(log, time.Now(), time.Now(), "dummy")
}

