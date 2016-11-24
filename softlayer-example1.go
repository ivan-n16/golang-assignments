package main

import (
        "fmt"
        "github.com/softlayer/softlayer-go/services"
        "github.com/softlayer/softlayer-go/session"
)

func main() {
        sess := session.New()
        sess.Debug = true

        service := services.GetAccountService(sess)

        vms, _ := service.Mask("id,hostname,domain,primaryIpAddress,primaryBackendIpAddress,accountId,billingItem[orderItem[order[userRecord[username]]]]").GetVirtualGuests()

        for _, vm := range vms {
                var primaryIpAddress = ""
                if vm.PrimaryIpAddress != nil {
                        primaryIpAddress = *vm.PrimaryIpAddress
                }

                fmt.Printf("ID: %d, Hostname: %s, Domain: %s, PublicIP: %s, PrivateIP: %s, AccountId: %d, Owner: %s\n", *vm.Id, *vm.Hostname, *vm.Domain, primaryIpAddress, *vm.PrimaryBackendIpAddress, *vm.AccountId, *vm.BillingItem.OrderItem.Order.UserRecord.Username)
        }
}
