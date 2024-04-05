package email

import (
	"context"
	"strings"
	"testing"
	"time"
)

const pwd = "******"
const username = "ihezebin@qq.com"

func TestInitEmail(t *testing.T) {
	_, err := NewClient(Config{
		Username: username,
		Password: pwd,
		Host:     HostQQMail,
		Port:     PortQQMail,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendEmail(t *testing.T) {
	client, err := NewClient(Config{
		Username: username,
		Password: pwd,
		Host:     HostQQMail,
		Port:     PortQQMail,
	})
	if err != nil {
		t.Fatal(err)
	}
	//now, err := timer.Parse("2006-01-02 15:04:05")
	msg := NewMessage().
		WithTitle("test").
		WithReceiver("86744316@qq.com").
		WithDate(time.Now()).
		WithSender("hezebin").
		WithHtml(`
			<html>
			<body>
				<h3 style="color:white;background-color:skyblue">
				"Hello World！This is a test mail！"
				</h3>
			</body>
			</html>
		`).
		WithAttach(NewAttach("test.txt", strings.NewReader("dsadsad")))
	err = client.Send(context.Background(), msg)
	if err != nil {
		t.Fatal("send mail err:", err)
	}

	t.Log("send mail successfully")
}
