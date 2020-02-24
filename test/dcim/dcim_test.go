package dcim_test

import (
    "fmt"
    runtimeclient "github.com/go-openapi/runtime/client"
    "github.com/hosting-de-labs/go-netbox/netbox/client"
    "os"
    "testing"
)

const (
    Host = "netbox.global.cloud.sap"
    DeviceID = 290
)

var Token = os.Getenv("TOKEN")

func TestDcim(t *testing.T) {
    tlsClient, _ := runtimeclient.TLSClient(runtimeclient.TLSClientOptions{InsecureSkipVerify: true})

    transport := runtimeclient.NewWithClient(Host, client.DefaultBasePath, []string{"https"}, tlsClient)
    transport.DefaultAuthentication = runtimeclient.APIKeyAuth("Authorization", "header", fmt.Sprintf("Token %v", Token))

    c := client.New(transport, nil)

    interfaceListTest := &InterfaceListTest{DeviceID:DeviceID, C:*c }
    if !t.Run("interfaceListTest", interfaceListTest.Run) {
        return
    }
}