package main

import (
	"credentity/generated"
	"errors"
	"fmt"
)

func main() {
	var tpvEnabled = generated.GatewayAttrsValues{
		Type:   generated.GatewayAttrs_TPV,
		Status: true,
	}
	var masterEnabled = generated.GatewayAttrsValues{
		Type:   generated.GatewayAttrs_SUPPORTS_MASTER,
		Status: true,
	}
	var g = &generated.Gateway{
		Name:               "upi",
		Status:             generated.GatewayStatus_ACTIVE,
		PaymentCodes:       []uint32{1001, 1002, 1003, 1004},
		BusinessCategories: []generated.BusinessCategories{generated.BusinessCategories_ECOMMERCE, generated.BusinessCategories_NONSKILL_GAMING},
		WhitelistMcc:       []uint32{4001, 4002, 4003},
		Attrs:              []generated.GatewayAttrs{generated.GatewayAttrs_SUPPORTS_MASTER, generated.GatewayAttrs_TPV},
	}

	var cred1 = &generated.GatewayCred{
		CredId:    "Cred_1",
		Status:    generated.CredStatus_CRED_ACTIVE,
		Gateway:   g,
		CredScope: generated.GatewayCredScope_DEDICATED,
		Attrs:     []*generated.GatewayAttrsValues{&tpvEnabled, &masterEnabled},
	}
	fmt.Println(g)
	fmt.Println("--------------")
	fmt.Println(cred1)
	fmt.Println("--------------")
	getTPVSupportedCred(g, cred1)
}

/*
questions
1. how do we ensure that gatewaycreds only have the attributes which are there in the gateway?
2.
*/
func getTPVSupportedCred(gateway *generated.Gateway, cred *generated.GatewayCred) {
	if contains(gateway.GetAttrs(), generated.GatewayAttrs_TPV) {
		val, err := attrValue(cred.Attrs, generated.GatewayAttrs_TPV)
		if err != nil {
			fmt.Println("TPV not supported by this cred")
		} else {
			fmt.Println("TPV is supported: ", val)
		}
	}
}

func contains(hay []generated.GatewayAttrs, needle generated.GatewayAttrs) bool {
	for _, v := range hay {
		if v == needle {
			return true
		}
	}
	return false
}

func attrValue(hay []*generated.GatewayAttrsValues, needle generated.GatewayAttrs) (*generated.GatewayAttrsValues, error) {
	for _, v := range hay {
		if v.Type == needle {
			return v, nil
		}
	}
	return nil, errors.New("attribute not found")
}
