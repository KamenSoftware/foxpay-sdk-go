package foxpay_sdk_go

import (
	"fmt"
	"github.com/DasenSoftware/foxpay-sdk-go/common"
	"github.com/DasenSoftware/foxpay-sdk-go/foxpay"
	"github.com/DasenSoftware/foxpay-sdk-go/util"
	"testing"
)

func TestFoxPay_GetBalance(t *testing.T) {
	fp := foxpay.NewFoxPay(
		foxpay.WithFoxPayObjAPPID("7IJNKYVX"),
		foxpay.WithFoxPayObjURL("http://139.159.184.46:7600"),
		foxpay.WithFoxPayObjPublicKey("MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCn4Q6mbmlSvH6kbeWJERz4rAQeTB/cEShKzgtkrWyaIZqHLgWh5iNdXyME0uaRUutFae1uc+J1yMyVU3cS+K36JUlqThmBHZ6/93KHsRvQ8FAcmBzKB2yVhW4qF0fA71yaWJzgNe93JI/4u3VSfu7tpy3ilPvmZlh6j9z9I+KKkQIDAQAB"),
		foxpay.WithFoxPayObjPrivateKey("MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAK7w5mS8/G7v7xjTAC5S+qx3E0OBFX6kt0kN3l9OHW7qA6cxFEP4DP04qRD3DCR4VvoHcNO0JwFv0Z3hc0afQt6qFGPcWMk9hFMTtB6gQwjOxLSRxDO2u1EVcfp13KzCLagx8/bfPM8+EMHkSXMPncIXmWXoy64r3aJmHOrCAgQnAgMBAAECgYAZjsMyC3Qbpvz41PatTd0mbh0H2ydvQZQvXZHTvZ9KMXEIL4Dk7yvCoND+VAFXoKcgTw76NtMOAC6RELtdIW5Mx+M4p8OZD+DeNqXUrXUpKd1odyagROSsqm03kwqDC+ZKIpU2f2SRST0HCi3ttXXMWqaKz3aeca6IaUPamoDGeQJBAPZHQRXlo1pbf0PJyjTgWkRB2FbZNCps9qoxhFMOOWf98rH/MvZ/Zi6YJMHqJssW+mK0lkMtgeHS5lwTgkGIqo8CQQC12MCRbBSdjL5hqUuFf8hUMQ+UC6yAEziiGAwkD5FhjpY07ylcywdJMH6srH4cFeYHTVfZPKLqk1yT46GheLjpAkAxSq+nL1ATnK6LJc836BOJB9jCATUkrKxuAf0nFni87KHvqFFN7s/H0aHBwhjDmzTAHr7YcTpGtYxvr2Pps+3XAkA7vhFN9X80X5fwh+ka2+dZ2aBvmAI9NZNmlZXvhvnRXkH09BnXtZAYOIl1e1oXKg6fmYZiBWzUukMxBxkD7qB5AkEAlP4x8e2red7uDNKk0iauppUFuZrX/dd3snm2ulAcC/qjacoXhCSF+KTApD9ScgZ7RJ4ZO2zr1lCfW5WbWaVi4Q=="),
	)
	fmt.Println(
		fp.GetBalance())
	fmt.Println(
		fp.GetApplicationOrder(common.OrderOrTradeNo{
			//OrderNo: util.StringPtr("004"),
			TradeNo: util.StringPtr("AP2023071310442022925526694"),
		}))
	fmt.Println(
		fp.CloseApplicationOrder(common.OrderOrTradeNo{
			OrderNo: util.StringPtr("004"),
		}))

	fmt.Println(
		fp.CreateApplicationOrder(common.OrderRequest{
			Subject:     "subjecttest",
			OrderNo:     "004",
			Amount:      "3.33",
			NotifyUrl:   util.StringPtr("nnotifyUrl"),
			RedirectUrl: util.StringPtr("redirectUrl"),
			TimeOut:     5000,
			Locale:      common.Zh_CN,
			Remark:      util.StringPtr("remarkTest"),
		}))
}