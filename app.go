package main

import (
  "flag"
  "fmt"
  "os"
  "os/user"
)

type Command struct {
  Username string
  Message string
}

func main() {
  cmd := Command{}

  user, err := user.Current()
  if err != nil {
    cmd.Username = ""
  } else {
    cmd.Username = user.Username
  }

  username := flag.String("name", fmt.Sprintf("%s",cmd.Username), "Username")
  message := flag.String("message", "", "Message (Reqired)")
  flag.Parse()

  cmd.Username = *username
  cmd.Message = *message

  if cmd.Message == "" {
    flag.PrintDefaults()
    os.Exit(1)
  }

  parse(cmd)
}

func parse(cmd Command){
  fmt.Printf("Hello %s, message: %s", cmd.Username, cmd.Message)
}
