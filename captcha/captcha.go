package captcha

import (
	"context"
	"log"

	"github.com/wzhanjun/dtclients/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client CaptchaClient

func init() {
	conn, err := grpc.NewClient(
		config.Cfg.Captcha.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("connect captcha client failed: %v", err)
	}
	defer conn.Close()
	client = NewCaptchaClient(conn)
}

func Send(account, drive, source string) (*SendCaptchaResponse, error) {
	if drive == "" {
		drive = "mobile"
	}
	return client.Send(context.Background(), &SendCaptchaRequest{Account: account, Driver: &drive, Source: &source})
}

func Verify(account, code, srouce string) (*VerifyCaptchaResponse, error) {
	return client.Verify(context.Background(), &VerifyCaptchaRequest{Account: account, Code: code, Source: &srouce})
}
