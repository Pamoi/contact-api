package main

import (
    "io"
    "net/http"
    "log"
    "time"
    "gopkg.in/gcfg.v1"
    "github.com/SlyMarbo/gmail"
)

type Config struct {
  Server struct {
    ServerName string
    PrivateKeyPath string
    PublicKeyPath string
  }
  Mail struct {
    Username string
    Password string
    Recipient string
  }
}

var conf Config

func sendMail(name, address, message string) bool {
  email := gmail.Compose("New message from " + name, "\nYou have a new message from " + name + " (" + address + ") on " + time.Now().Format("Jan 2 15:04:05") + ":\n\n"  + message)
  email.From = conf.Mail.Username
  email.Password = conf.Mail.Password
  email.AddRecipient(conf.Mail.Recipient)

  err := email.Send()
  if err != nil {
      log.Println("Failed to send email: ", err)
      return false
  }

  return true
}

func PostMessage(w http.ResponseWriter, req *http.Request) {
  name := req.FormValue("name")
  email := req.FormValue("email")
  message := req.FormValue("message")

  if name == "" || email == "" || message == "" {
    w.WriteHeader(http.StatusUnprocessableEntity)
    io.WriteString(w, "{\"status\": \"Missing aruments.\"}")
    return
  }

  if sendMail(name, email, message) {
    io.WriteString(w, "{\"status\": \"Success.\"}")
  } else {
    w.WriteHeader(http.StatusInternalServerError)
    io.WriteString(w, "{\"status\": \"Failed to send email.\"}")
  }
}

func main() {
  err := gcfg.ReadFileInto(&conf, "config.gcfg")
  if err != nil {
    log.Fatal("Failed to parse config file: ", err)
  }

  http.HandleFunc("/message", PostMessage)
  err = http.ListenAndServeTLS(conf.Server.ServerName, conf.Server.PublicKeyPath, conf.Server.PrivateKeyPath, nil)
  if err != nil {
    log.Fatal("Failed to start HTTPS server: ", err)
  }
}
