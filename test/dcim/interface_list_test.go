package dcim_test

import (
    "fmt"
    "github.com/hosting-de-labs/go-netbox/netbox/client"
    "github.com/hosting-de-labs/go-netbox/netbox/client/dcim"
    "github.com/stretchr/testify/assert"
    "reflect"

    "testing"
)

type InterfaceListTest struct {
    DeviceID int64
    C   client.NetBox
}

func (k *InterfaceListTest) Run(t *testing.T) {
    _ = t.Run("byDeviceID", k.ListInterfacesByDeviceID)
}

func (k *InterfaceListTest) ListInterfacesByDeviceID(t *testing.T) {
    params := dcim.NewDcimInterfacesListParams()
    params.SetDeviceID(&k.DeviceID)
    //We want only type LAG so pass in 200
    typ := "lag"
    params.SetType(&typ)
    res, err := k.C.Dcim.DcimInterfacesList(params, nil)
    if err != nil {
        fmt.Printf("err DcimInterfacesList: %v", err)
        fmt.Printf("payload: %v", res)
    }
    fmt.Printf("DEBUG %v", res)
    assert.NoError(t, err)
    assert.True(t, reflect.TypeOf(res.Payload.Results[0].Type.Value).String() == "*string")
}
