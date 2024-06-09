package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)



func main(){

  listner, err := net.Listen("tcp",":8000")
  logrus.Info("Server is running")
  if err != nil {
    panic(err)
  } 

  for {
    conn,err := listner.Accept()
    
    if err != nil{
      panic(err)
    }

    go handleConnection(conn)
  } 

}


func handleConnection(conn net.Conn){
  defer conn.Close()
  logrus.Info("New Connection")
  for {
     buf := make([]byte, 1024)

     // Read the incoming connection into the buffer.
     reqLen, err := conn.Read(buf)
     if err != nil {
        logrus.Error("Connection Lost")
    }

     // Print the message to the console. 
     fmt.Println("Received data:", string(buf[:reqLen]))

  }

}


