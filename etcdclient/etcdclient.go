// based on: https://github.com/octoblu/go-simple-etcd-client
package etcdclient

import (
  "time"
  "context"
  "github.com/coreos/etcd/client"
)

type Client interface {
  Ls(dir string, recursive, leafonly bool) ([]string, error)
  Get(key string) (string, error)
  Set(key, val string, directory bool) error
  Del(key string) error
}

type SimpleClient struct {
  etcd client.Client
}

func NewClient(uri string) (Client, error) {
  etcd, err := client.New(client.Config{
    Endpoints: []string{uri},
    Transport: client.DefaultTransport,
    HeaderTimeoutPerRequest: 5 * time.Second,
  })
  if err != nil {
    return nil, err
  }
  return &SimpleClient{etcd}, nil
}

func (cli *SimpleClient) Ls(dir string, recursive, leafonly bool) ([]string, error) {
  api := client.NewKeysAPI(cli.etcd)
  options := &client.GetOptions{Sort: true, Recursive: recursive}
  res, err := api.Get(context.Background(), dir, options)
  if err != nil {
    if client.IsKeyNotFound(err) {
      return make([]string, 0), nil
    }
    return make([]string, 0), err
  }
  if leafonly {
    return nodesToStringSlice(extractLeafNodes(res.Node.Nodes)), nil
  }
  return nodesToStringSlice(res.Node.Nodes), nil
}

func (cli *SimpleClient) Get(key string) (string, error) {
  api := client.NewKeysAPI(cli.etcd)
  res, err := api.Get(context.Background(), key, nil)
  if err != nil {
    if client.IsKeyNotFound(err) {
      return "", nil
    }
    return "", err
  }
  return res.Node.Value, nil
}

func (cli *SimpleClient) Set(key, val string, directory bool) error {
  api := client.NewKeysAPI(cli.etcd)
  options := &client.SetOptions{Dir: directory}
  _, err := api.Set(context.Background(), key, val, options)
  return err
}

func (cli *SimpleClient) Del(key string) error {
  api := client.NewKeysAPI(cli.etcd)
  _, err := api.Delete(context.Background(), key, nil)
  if err != nil {
    if client.IsKeyNotFound(err) {
      return nil
    }
  }
  return err
}

func extractLeafNodes(nodes client.Nodes) client.Nodes {
  var leafnodes client.Nodes
  for _, node := range nodes {
    if ! node.Dir {
      leafnodes = append(leafnodes, node)
    }
  }
  return leafnodes
}

func nodesToStringSlice(nodes client.Nodes) []string {
  var keys []string
  for _, node := range nodes {
    keys = append(keys, node.Key)
    for _, key := range nodesToStringSlice(node.Nodes) {
      keys = append(keys, key)
    }
  }
  return keys
}
