package main

import (
  "log"
  "time"
  "context"
  "github.com/coreos/etcd/client"
)

func put(api client.KeysAPI, key string, val string){
	resp, err := api.Set(context.Background(), key, val, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
}

func get(api client.KeysAPI, key string){
	resp, err := api.Get(context.Background(), key, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Get is done. Metadata is %q\n", resp)
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
}

func delete(api client.KeysAPI, key string){
	_, err := api.Delete(context.Background(), key, &client.DeleteOptions{})
	if err != nil {
		log.Fatal(err)
	}
}

func testdata_injection(){
  cli, err := client.New(client.Config{
    Endpoints: []string{"http://127.0.0.1:62379"},
    Transport: client.DefaultTransport,
    HeaderTimeoutPerRequest: 5 * time.Second,
  })
  if err != nil {
    log.Fatal(err)
  }
  api := client.NewKeysAPI(cli)
  put(api, "/jp/ac/titech/e/ict/net/.SOA", "net.ict.e.titech.ac.jp.\tnet-root.net.ict.e.titech.ac.jp.\t3600\t600\t86400\t10")
  put(api, "/jp/ac/titech/e/ict/net/.NS/ns1", "ns1.net.ict.e.titech.ac.jp.")
  put(api, "/jp/ac/titech/e/ict/net/.NS/ns2", "ns2.net.ict.e.titech.ac.jp.")
  put(api, "/jp/ac/titech/e/ict/net/.NS/ns1.ttl", "3600")
  put(api, "/jp/ac/titech/e/ict/net/.NS/ns2.ttl", "3600")
  put(api, "/jp/ac/titech/e/ict/net/ns1/.A", "131.112.21.100")
  put(api, "/jp/ac/titech/e/ict/net/ns2/.A", "131.112.21.101")
  put(api, "/jp/ac/titech/e/ict/net/ns1/.A.ttl", "3600")
  put(api, "/jp/ac/titech/e/ict/net/ns2/.A.ttl", "3600")
  put(api, "/jp/ac/titech/e/ict/net/www/.CNAME", "yamaoka-kitaguchi-lab.github.io")
  put(api, "/jp/ac/titech/e/ict/net/.CNAME.ttl", "60")
  put(api, "/jp/ac/titech/e/ict/net/www2/.A", "131.112.21.120")
  put(api, "/jp/ac/titech/e/ict/net/www2/.A.ttl", "360")
  put(api, "/jp/ac/titech/e/ict/net/wiki/.A/1st", "131.112.21.131")
  put(api, "/jp/ac/titech/e/ict/net/wiki/.A/2nd", "131.112.21.132")
  put(api, "/jp/ac/titech/e/ict/net/wiki/.A/1st.ttl", "120")
  put(api, "/jp/ac/titech/e/ict/net/wiki/.A/2nd.ttl", "120")
  put(api, "/jp/ac/titech/e/ict/net/dns-acme/.TXT", "challenge")
  put(api, "/jp/ac/titech/e/ict/net/dns-acme/.TXT.ttl", "1200")
  put(api, "/jp/ac/titech/e/ict/net/.MX/1st", "10\tfilter1.nap.gsic.titech.ac.jp.")
  put(api, "/jp/ac/titech/e/ict/net/.MX/1st.ttl", "3600")
  put(api, "/jp/ac/titech/e/ict/net/.MX/2nd", "15\tfilter2.nap.gsic.titech.ac.jp.")
  put(api, "/jp/ac/titech/e/ict/net/.MX/2nd.ttl", "3600")
}

func main(){
  // testdata_injection()
}
