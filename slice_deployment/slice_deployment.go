package main

import (
  "fmt"
  "github.com/Adnei/kube_slice/k8s_cluster"
)

func main(){
  message := k8s_cluster.Hello("Adnei")
  fmt.Println(message)
}
