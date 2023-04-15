package domain

import (
	"fmt"
	"testing"
)

func TestAddress_SigVerify_OK(t *testing.T) {
	msg := "WalletConnectログイン\n        \n        Nonce: MmXe9WOJ"
	target := Address("0x45d0BF1e8dcFe508aD1024A6fEdB2a21f9FAbE68")
	signature := "0x3a9cf0d4b2f2a201911eef1b56633a24373c404600fc409629c98d493bb693297bc163b2687a9f94dda6472dcaf1dd1cd5943dda43cffc1fc23dbb9d2f4fe4bb1b"
	if err := target.SigVerify(msg, signature); err != nil {
		t.Errorf("<%s> verify failed. err:%+v\n", t.Name(), err)
	}
	fmt.Printf("<%s> end.\n\n", t.Name())
}

func TestAddress_SigVerify_Nonce_Unmatched(t *testing.T) {
	msg := "Sign\n        \n        Nonce: FtqDQfmbX8y"
	target := Address("0x45d0BF1e8dcFe508aD1024A6fEdB2a21f9FAbE68")
	signature := "0x3a9cf0d4b2f2a201911eef1b56633a24373c404600fc409629c98d493bb693297bc163b2687a9f94dda6472dcaf1dd1cd5943dda43cffc1fc23dbb9d2f4fe4bb1b"
	if err := target.SigVerify(msg, signature); err == nil {
		t.Errorf("<%s> verify check failed. err:%+v\n", t.Name(), err)
	}
	fmt.Printf("<%s> end.\n\n", t.Name())
}
