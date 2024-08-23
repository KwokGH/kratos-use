package sms

import (
	"context"
	"gopkg.in/gomail.v2"
	"strings"
)

// SendEmailCode 发送邮箱验证码

type SendEmailCodeInput struct {
	Email    string
	From     string
	Subject  string
	Smtp     string
	Port     int32
	AuthCode string
	// 为空使用默认的模板
	Template string
	Code     string
}

func SendEmailCode(ctx context.Context, input *SendEmailCodeInput) error {
	tpl := EmailBodyTemplate
	if input.Template != "" {
		tpl = input.Template
	}

	body := strings.NewReplacer("{:code}", input.Code).Replace(tpl)
	m := gomail.NewMessage()
	m.SetHeader("From", input.From)
	m.SetHeader("To", input.Email)
	m.SetHeader("Subject", input.Subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(input.Smtp, int(input.Port), input.From, input.AuthCode)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

const EmailBodyTemplate = `
<!DOCTYPE html>
<html>

<head>
    <style>
        body {
            /* background-color: #EFEFEF; */
            display: flex;
            /* height: 100vh; */
            justify-content: center;
            align-items: center;
        }

        .message {
            /* background-color: #EFEFEF; */
            color: #000;
            text-align: center;
            width: 100%;
            height: 100%;
            padding: 6vw;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
        }

        .page-title {
            display: flex;
            justify-content: center;
            align-items: center;
            overflow: hidden;
            border-radius: 10px;
            width: 15vw;
            height: 15vw;
            margin-bottom: 3vw;
        }

        .page-body {
            width: 100%;
            height: 40%;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            background-color: #fff;
            border-radius: 10px;
        }


        .body-title {
            width: 70%;
            /* font-size: 28px; */
            text-align: center;
        }

        .page-content {
            width: 70%;
            /* font-size: 22px; */
            text-align: justify;
        }

        .verify-code {
            width: 30%;
            /* font-size: 28px; */
        }

        .page-info {
            width: 50%;
            /* font-size: 22px; */
            text-align: center;
        }
    </style>
</head>

<body>
    <div class="message">
        <div class="page-title">
            <img width="100%" height="100%"
                src="https://prod-ayayi.tos-ap-southeast-1.volces.com/website/ayayi-logo2.svg" alt="AYAYI" />
        </div>
        <div class="page-body">
            <div class="body-title">
                <h2>Verify your email address with AYAYI</h2>
            </div>
            <div class="page-content">
                <p>
                    You have requested to sign up to AYAYI. Please enter the following verification code to continue the
                    signup process.
                </p>
            </div>
            <div class="verify-code">
                <h2>{:code}</h2>
            </div>
            <div class="page-info">
                <p>The verification code will expire in 5 minutes.</p>
            </div>
        </div>

    </div>
</body>

</html>
`
