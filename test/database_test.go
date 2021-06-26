package test

import (
    "context"
    "github.com/sia-fl/medicine/src/proto"
    "github.com/sia-fl/medicine/src/util"
    "testing"
)

func TestIds(t *testing.T) {
    Injection(func(
        idsRpcSer proto.IdInMapServicesClient,
    ) {
        _, err := idsRpcSer.GetIdInMap(context.Background(), &proto.GetIdInMapReq{EId: 1, Name: "map_ids"})
        util.CE(err)
    })
}
