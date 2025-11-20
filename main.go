package main

import (
    "bytes"
    "fmt"
    "log"
    "os"
    "log/slog"
)

// The go standard library provides straightforward tools for outputting logs
// from Go programs, with the log package for free-form output and the log/slog
// package for structured output.

func main() {

    // Println makes use of the standard logger.
    log.Println("standard logger")
    
    // Loggers can be configured with flags to set their output format. By default
    // the standard logger has the log.Ldate and log.Ltime flags set, and these are
    // collected in log.LstdFlags.
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)
    log.Println("with micro")

    // You can also emit the file name and line from which the log function is called.
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("with file/line")

    // You can also create a custom logger and pass it around. When doing so, you can
    // set a prefix to distinguish its output from other loggers.
    mylog := log.New(os.Stdout, "my:", log.LstdFlags)
    mylog.Println("from mylog")

    // We can also set a prefix for existing loggers.
    mylog.SetPrefix("ohmy:")
    mylog.Println("from mylog")

    // Loggers can have custom output argets; any io.Writer works.
    var buf bytes.Buffer
    buflog := log.New(&buf, "buf:", log.LstdFlags)

    // This call writes the log output into buf.
    buflog.Println("hello")

    // This will actually show on standard output.
    fmt.Print("from buflog:", buf.String())

    // In addition to the message, slog output can contain an arbitrary number
    // of key=value pairs.
    jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
    myslog := slog.New(jsonHandler)
    myslog.Info("hi there")
    
    // Sample output; the date and time emitted will depend on when the
    // example run.
    myslog.Info("hello again", "key", "val", "age", 25)
}