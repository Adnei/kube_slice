package kube_slice

import (
  "github.com/Adnei/k8s_cluster"
)

func HelloAdnei() string {
  return k8s_cluster.Hello("Adnei");
}
