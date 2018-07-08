package main

import (
  "log"
  "./etcdclient"
  "./discoapi"
)

func test_myetcdclient_2(){
  cli, err := etcdclient.NewClient("http://127.0.0.1:62379")
  if err != nil {
    log.Fatal(err)
  }

  // can resolve
  cli.Set("/jp/ac/titech/e/ict/net/eee6/.A", "2.2.2.2", false)
  cli.Set("/jp/ac/titech/e/ict/net/eee6/.A.ttl", "30", false)
  cli.Set("/jp/ac/titech/e/ict/net/eee6/.A.enabled", "1", false)
  cli.Set("/jp/ac/titech/e/ict/net/eee6/.A.comment", "test", false)
  cli.Set("/jp/ac/titech/e/ict/net/eee6/.A.uuid", "1234567", false)

  api, err := discoapi.NewDiscoAPI("http://127.0.0.1:62379")
  if err != nil {
    log.Fatal(err)
  }

  err = api.Create(&discoapi.RR{
    Type: "A",
    FQDN: "qqq7.net.ict.e.titech.ac.jp.",
    Value: "1.1.1.1",
    TTL: 300,
    Enabled: true,
    Comment: "samplecase",
    UUID: "1234567890",
  })
  if err != nil {
    log.Fatal(err)
  }

  res, _ := cli.Ls("/", true, false)
  log.Printf("%q", res)
}

func main(){
  test_myetcdclient_2()
}
