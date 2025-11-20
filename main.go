package main

import (
    "errors"
    "fmt"
    "io"
    "os/exec"
)

// There are times Go programs will need to spawn other processes.

func main() {

    // We'll start with a simple command that takes no arguments or input
    // and just prints something to stdout.
    dateCmd := exec.Command("date")

    // The Output method runs the command, waits for it to finish, and collects
    // its standard output. If there are no errors, dateOut will hold bytes with
    // the date info.
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }

    fmt.Println("> date")
    fmt.Println(string(dateOut))

    // Output and other methods of Command will return *exec.Error if there was
    // a problem executing the command and *exec.ExitError if the command di run
    // but exited with a non-zero return code.
    _, err = exec.Command("date", "-x").Output()
    if err != nil {
        var execErr *exec.Error
        var exitErr *exec.ExitError
        switch {
        case errors.As(err, &execErr):
            fmt.Println("failed executing:", err)
        case errors.As(err, &exitErr):
            exitCode := exitErr.ExitCode()
            fmt.Println("command exit rc =", exitCode)
        default:
            panic(err)
        }
    }

    // Next we will look at a slightly more involved case where we pipe data to the
    // external process on its stdin and collect the results from its stdout.
    grepCmd := exec.Command("grep", "hello")

    // Here we explicitly grab input/output pipes, start the process, write some
    // input to it, read the resulting output, and finally wait for the process to exit.
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := io.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    // Note: when spawning commands you must provide an explicitly delineated
    // command and argument array, rather than passing in a one command string.
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}