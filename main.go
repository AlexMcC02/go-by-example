package main

import (
    "os"
    "os/exec"
    "syscall"
)

// There may be cases we want to replace the current Go process with a new one,
// to do this we use go's implementation of the exec function.

func main() {

    // Go requires an absolute path to the binary we want to execute, so we use
    // exec.LookPath to find it.
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    // Exec requires arguments in slice form (as opposed to one big string).
    // We'll give ls a few common arguments.
    // Note: the first argument should be the program name.
    args := []string{"ls", "-a", "-l", "-h"}
    
    // Exec also requires certain environment variables.
    env := os.Environ()

    // Here's the actual syscall.Exec call. If successfuly, the execution of
    // our process will end and be replaced by the newly created process.
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}