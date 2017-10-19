package main

import (
  "os/exec"
  "os"
  "fmt"
  "log"
  "strings"
  "errors"
  "io/ioutil"
)

func extractDescription (rawBranchName string) string {
  arr := strings.Split(rawBranchName, "-")

  if len(arr) < 2 {
    return rawBranchName
  }

  slice := arr[0:2]
  branch := strings.Join(slice, "-")
  return branch
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func writeMsg(file string, msg []byte) {
  err := ioutil.WriteFile("./" + file, msg, 0644)
  check(err)
}

func readMsg(file string) []byte {
  msg, err := ioutil.ReadFile("./" + file)
  check(err)

  return msg
}

func prepareMessage(rawBranch string) string {
  var branchName string
  commitMsgFile := os.Args[1]
  msg := readMsg(commitMsgFile)

  switch rawBranch {
  case "master":
    branchName = ""
    break
  default:
    arr := strings.Split(rawBranch, "/") // feature/HS-

    if len(arr) < 2 {
      log.Fatal(errors.New("Not support branches without type "))
    }

    branchName = arr[1]
  }

  return fmt.Sprintf("%s %s", strings.TrimSpace(branchName), strings.TrimSpace(string(msg)))
}

func main() {
  cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD") // exec branch command
  outBytes, err := cmd.Output()

  if err != nil {
    log.Fatal(err)
  }

  outputResult := extractDescription(strings.TrimSpace(string(outBytes)))

  msg := []byte(prepareMessage(outputResult))
  writeMsg(os.Args[1], msg)
}

