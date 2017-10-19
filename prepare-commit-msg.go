package main

import (
  "os/exec"
  "os"
  "log"
  "strings"
  "./helpers"
)

func main() {
  cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD") // exec branch command
  outBytes, err := cmd.Output()

  if err != nil {
    log.Fatal(err)
  }

  outputResult := helpers.ExtractDescription(strings.TrimSpace(string(outBytes)))

  msg := []byte(helpers.PrepareMessage(outputResult))
  helpers.WriteMsg(os.Args[1], msg)
}

