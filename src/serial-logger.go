package main

import (
    "fmt"
    "os"
    "time"
    "github.com/tarm/serial"
)

const BUFSIZE = 1024

var buf  []byte = make([]byte, BUFSIZE)
var bufi   int  = 0

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
    
    // service loop
    for {
        // read line
        var t0 time.Time = time.Now()
        for {
            // blocking read
            n, err := s.Read(buf[bufi:bufi+1])
            if err != nil {
                fmt.Println(err)
                os.Exit(4)
            }
            
            // guard: disconnect case?
            if n==0 {
                fmt.Println("n==0")
                os.Exit(5)
            }
            
            bufi++
            
            // check exit condition
            if buf[bufi-1]=='\n' {
                break
            }
        }
        var t1 time.Time = time.Now()
        
        // process line
        append2log(log, t0, t1, string(buf[0:bufi]))
        
        // cleanup
        bufi = 0
    }
}

