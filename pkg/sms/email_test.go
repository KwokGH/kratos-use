package sms

import (
	"context"
	"testing"
)

func TestSendEmailCode(t *testing.T) {
	input := &SendEmailCodeInput{
		Email:    "",
		From:     "",
		Subject:  "",
		Smtp:     "",
		Port:     25,
		AuthCode: "",
		Template: EmailBodyTemplate,
	}

	err := SendEmailCode(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("发送成功")
}
