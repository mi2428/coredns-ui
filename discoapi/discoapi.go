package discoapi

import (
  "strings"
  "../etcdclient"
)

type RR struct {
  Type string
  FQDN string
  Value string
  Priority uint
  TTL uint
  Enabled bool
  Comment string
  UUID string
}

type API interface {
  Create(rr *RR) error
}

type DiscoAPI struct {
  etcdcli etcdclient.Client
}

func NewDiscoAPI(etcduri string) (API, error) {
  etcdcli, err := etcdclient.NewClient(etcduri)
  if err != nil {
    return nil, err
  }
  return &DiscoAPI{etcdcli}, nil
}

func (dapi *DiscoAPI) Create(rr *RR) error {
  etcd := dapi.etcdcli
  rrdir := fqdnToKeyFormat(rr.FQDN) + "/." + rr.Type
  rrkey := rrdir + "/" + rr.UUID
  rrkey_ttl := rrkey + ".ttl"
  rrkey_enabled := rrkey + ".enabled"
  rrkey_comment := rrkey + ".comment"
  rrkey_uuid := rrkey + ".uuid"

  rrvalue := rr.Value
  if rr.Type == "MX" {
    rrvalue = string(rr.Priority) + "\t" + rr.Value
  }

  rrenabled := "1"
  if ! rr.Enabled {
    rrenabled = "0"
  }

  err := etcd.Set(rrdir, "", true)
  if err != nil {
    return err
  }
  err = etcd.Set(rrkey, rrvalue, false)
  if err != nil {
    return err
  }
  err = etcd.Set(rrkey_ttl, string(rr.TTL), false)
  if err != nil {
    return err
  }
  err = etcd.Set(rrkey_enabled, rrenabled, false)
  if err != nil {
    return err
  }
  err = etcd.Set(rrkey_comment, rr.Comment, false)
  if err != nil {
    return err
  }
  err = etcd.Set(rrkey_uuid, rr.UUID, false)
  if err != nil {
    return err
  }
  return nil
}

func (dapi *DiscoAPI) Read() ([]RR, error) {
  return nil, nil
}

func fqdnToKeyFormat(fqdn string) string {
  fa := strings.Split(fqdn, ".")
  var etcdkey []string
  for i := len(fa) - 1; i >= 0; i-- {
    etcdkey = append(etcdkey, fa[i])
  }
  return strings.Join(etcdkey, "/")
}
