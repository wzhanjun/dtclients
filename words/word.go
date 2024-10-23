package words

import (
	context "context"
	"log"

	"github.com/wzhanjun/dtclients/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client WordsClient

func init() {
	conn, err := grpc.NewClient(
		config.Cfg.Words.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("connect words client failed: %v", err)
	}
	// defer conn.Close()
	client = NewWordsClient(conn)
}

func CheckSensitive(text string) (*CheckSensitiveResponse, error) {
	return client.CheckSensitive(context.Background(), &CheckSensitiveRequest{Text: text})
}
