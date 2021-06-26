package cfg

import (
    "fmt"
    "github.com/sia-fl/medicine/src/proto"
    "google.golang.org/grpc"
)

type Database struct {
    Drive  string
    Conn   string
    Prefix string
}

type ServiceIds struct {
    Host string
    Port string
}

func (s *ServiceIds) GenServiceIdsClient() proto.IdInMapServicesClient {
    conn, err := grpc.Dial(fmt.Sprintf("%s:%s", s.Host, s.Port), grpc.WithInsecure())
    if err != nil {
        panic(err)
    }
    return proto.NewIdInMapServicesClient(conn)
}
