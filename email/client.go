package email

import (
	"context"
	"fmt"

	"gopkg.in/gomail.v2"
)

type Client struct {
	config Config
	kernel *gomail.Dialer
}

func NewClient(config Config) (*Client, error) {
	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	closer, err := dialer.Dial()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = closer.Close()
	}()
	return &Client{kernel: dialer, config: config}, nil
}

func (client *Client) Kernel() *gomail.Dialer {
	return client.kernel
}

func (client *Client) Send(ctx context.Context, message Message) error {
	username := client.config.Username
	if message.Sender == "" {
		message.Sender = username
	} else {
		message.Sender = fmt.Sprintf(`%s <%s>`, message.Sender, username)
	}
	return client.kernel.DialAndSend(message.transport())
}
