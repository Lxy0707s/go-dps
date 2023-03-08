package template_dp

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	o := &MsgTmp{
		Message: &Sms{name: "default"},
	}
	info := &MsgInfo{
		Name:    "test-sms",
		Address: "sms-center",
	}
	err := o.genAndSendOTP("", info)
	if err != nil {
		return
	}

	fmt.Println()

	o = &MsgTmp{
		Message: &Email{name: "post-tmp"},
	}
	info = &MsgInfo{
		Name:    "user-email",
		Address: "email-center",
	}
	err = o.genAndSendOTP("", info)
	if err != nil {
		return
	}
}
