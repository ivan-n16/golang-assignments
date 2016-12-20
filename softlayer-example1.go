package main

import (
        "fmt"
        "github.com/softlayer/softlayer-go/services"
        "github.com/softlayer/softlayer-go/session"
        "github.com/softlayer/softlayer-go/sl"
)

func main() {
        sess := session.New()
        sess.Debug = true

        service := services.GetAccountService(sess)

        vms, _ := service.Mask("id,hostname,domain,primaryIpAddress,primaryBackendIpAddress,accountId,billingItem[orderItem[order[userRecord[username]]]]").GetVirtualGuests()

        for _, vm := range vms {
/* Redundant check for nil
                var primaryIpAddress = ""
                if vm.PrimaryIpAddress != nil {
                        primaryIpAddress = *vm.PrimaryIpAddress
                }
*/
                fmt.Printf("ID: %d, Hostname: %s, Domain: %s, PublicIP: %s, PrivateIP: %s, AccountId: %d, Owner: %s\n", *vm.Id, *vm.Hostname, *vm.Domain, sl.Grab(vm, "PrimaryIpAddress"), *vm.PrimaryBackendIpAddress, *vm.AccountId, *vm.BillingItem.OrderItem.Order.UserRecord.Username)
        }
}
