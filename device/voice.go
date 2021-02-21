package main

import (
    "os/exec"
    "log"
)

func main() {
    s := "Make the computer speak"
    cmd := exec.Command("espeak", s)
    if err := cmd.Run(); err != nil {
        log.Fatal(err)
    }
}

