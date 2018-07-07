package main

import (
  "fmt"
  "log"
  "time"
  "context"
  "github.com/coreos/etcd/client"
  "github.com/coreos/etcd/clientv3"
)

func client_demo() {
  cfg := client.Config{
    Endpoints: []string{"http://127.0.0.1:62379"},
    Transport: client.DefaultTransport,
    HeaderTimeoutPerRequest: time.Second,
  }
  c, err := client.New(cfg)
  if err != nil {
    log.Fatal(err)
  }
  kapi := client.NewKeysAPI(c)
  log.Print("Setting '/foo' key with 'bar' value")
	resp, err := kapi.Set(context.Background(), "/foo2", "bar", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
	// get "/foo" key's value
	log.Print("Getting '/foo' key value")
	resp, err = kapi.Get(context.Background(), "/foo2", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
}

func client_v3_demo(){
  cli, err := clientv3.New(clientv3.Config{
    Endpoints: []string{"localhost:62379"},
    DialTimeout: 5 * time.Second,
  })
  if err != nil {
    log.Fatal(err)
  }
  defer cli.Close()

  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  _, err = cli.Put(ctx, "foo", "bar")
  if err != nil {
    log.Fatal(err)
  }

  resp, err := cli.Get(ctx, "foo")
  if err != nil {
    log.Fatal(err)
  }
  for _, ev := range resp.Kvs {
    fmt.Printf("%s : %s\n", ev.Key, ev.Value)
  }

  // count keys about to be deleted
  gresp, err := cli.Get(ctx, "foo", clientv3.WithPrefix())
  if err != nil {
      log.Fatal(err)
  }

  // delete the keys
  dresp, err := cli.Delete(ctx, "foo", clientv3.WithPrefix())
  if err != nil {
      log.Fatal(err)
  }

  fmt.Println("Deleted all keys:", int64(len(gresp.Kvs)) == dresp.Deleted)

  resp, err = cli.Get(ctx, "foo")
  if err != nil {
    log.Fatal(err)
  }
  for _, ev := range resp.Kvs {
    fmt.Printf("%s : %s\n", ev.Key, ev.Value)
  }
}

func putv3(cli *clientv3.Client, key string, val string){
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()
  _, err := cli.Put(ctx, key, val)
  if err != nil {
    log.Fatal(err)
  }
}

func getv3(cli *clientv3.Client, key string){
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()
  resp, err := cli.Get(ctx, key)
  if err != nil {
    log.Fatal(err)
  }
  for _, ev := range resp.Kvs {
    fmt.Printf("%s : %s\n", ev.Key, ev.Value)
  }
}

func demov3() {
  cli, err := clientv3.New(clientv3.Config{
    Endpoints: []string{"localhost:62379"},
    DialTimeout: 5 * time.Second,
  })
  if err != nil {
    log.Fatal(err)
  }
  defer cli.Close()
  putv3(cli, "/jp/ac/titech/e/ict/net/www/.A", "131.112.21.93")
  putv3(cli, "/jp/ac/titech/e/ict/net/www/.A.ttl", "61")
  getv3(cli, "/jp/ac/titech/e/ict/net/www/.A")
  getv3(cli, "/jp/ac/titech/e/ict/net/www/.A.ttl")
}

func putv2(api client.KeysAPI, key string, val string){
	resp, err := api.Set(context.Background(), key, val, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
}

func getv2(api client.KeysAPI, key string){
	resp, err := api.Get(context.Background(), key, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Get is done. Metadata is %q\n", resp)
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
}

func deletev2(api client.KeysAPI, key string){
	_, err := api.Delete(context.Background(), key, &client.DeleteOptions{})
	if err != nil {
		log.Fatal(err)
	}
}

func demov2(){
  cli, err := client.New(client.Config{
    Endpoints: []string{"http://127.0.0.1:62379"},
    Transport: client.DefaultTransport,
    HeaderTimeoutPerRequest: 5 * time.Second,
  })
  if err != nil {
    log.Fatal(err)
  }
  api := client.NewKeysAPI(cli)
  putv2(api, "/jp/ac/titech/e/ict/net/www/.A", "131.112.21.93")
  putv2(api, "/jp/ac/titech/e/ict/net/www/.A.ttl", "61")
  getv2(api, "/jp/ac/titech/e/ict/net/www/.A")
  getv2(api, "/jp/ac/titech/e/ict/net/www/.A.ttl")
  // deletev2(api, "/jp/ac/titech/e/ict/net/www/.A")
  // deletev2(api, "/jp/ac/titech/e/ict/net/www/.A.ttl")
}



func main(){
  demov2()
}
