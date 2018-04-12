// +build examples
package examples

import (
	"fmt"
	"github.com/ignw/cisco-aci-go-sdk/src/service"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

func main() {
	var client *service.Client
	var tenant *models.Tenant
	var err error

	client = service.GetClient()

	tenant = client.Tenants.New("Example-Tenant", "This is an example tenant")

	err = client.Tenants.Save(tenant)

	if err != nil {
		fmt.Printf("Error creating tenant: %s", err.Error())
	} else {
		fmt.Println("Successfully created new tenant!")
	}
}