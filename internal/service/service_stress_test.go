package service

import (
	"context"
	"fmt"
	"github.com/ChungKiet/cake-interview/api/server"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
const letterNums = "0123456789"

func NewServer() (server.CAKEServiceClient, error) {
	time_out := 3 // TODO: TBU
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(time_out)*60*time.Second)
	maxMsgSize := 1024 * 1024 * 1024 //1GB
	//
	transportOption := grpc.WithInsecure()
	conn, err := grpc.DialContext(
		ctx,
		"localhost:9005", // TODO: TBU
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize),
			grpc.MaxCallSendMsgSize(maxMsgSize)),
		transportOption,
	)
	if err != nil {
		return nil, err
	}

	client := server.NewCAKEServiceClient(conn)
	return client, nil
}

func TestCAKEServiceService_Register(t *testing.T) {
	concurrentUser := 100
	logger := log.NewHelper(log.DefaultLogger)
	cli, err := NewServer()
	if err != nil {
		logger.Error("Init server error", "err", err)
		panic(err)
	}
	//ctx := context.TODO()
	emailPostFix := "@gmail.com"
	// assume that all user use their user_name to register new account
	wg := sync.WaitGroup{}
	totalErr := 0
	for i := 0; i < concurrentUser; i++ {
		wg.Add(1)
		ctxWithTimeOut, _ := context.WithTimeout(context.Background(), 5*time.Second)
		go func(ctx context.Context) {
			userName := RandStringBytes(20)
			_, err = cli.Register(ctx, &server.RegisterNewAccountRequest{
				FullName:    RandStringBytes(15),
				PhoneNumber: RandPhoneNumber(),
				Account:     userName,
				UserName:    userName,
				Email:       RandStringBytes(10) + emailPostFix,
				Password:    RandStringBytes(10),
				Birthday:    time.Now().Unix(),
			})

			if err != nil {
				totalErr += 1
			}
		}(ctxWithTimeOut)
	}
	wg.Wait()
	fmt.Printf("Failed case / Total case: %d / %d", totalErr, concurrentUser)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandPhoneNumber() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterNums))]
	}
	return string(b)
}
