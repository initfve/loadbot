package cli

import (
	"context"
	"fmt"

	"github.com/kuzxnia/loadbot/lbot/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// checks if process is running in local system

// tutaj nie powinno wchodzić proto
func StartDriver(conn grpc.ClientConnInterface, request *proto.StartRequest) (err error) {
	// todo: mapowanie to proto
	log.Info("🚀 Starting stress test")

	client := proto.NewStartProcessClient(conn)

	reply, err := client.Run(context.TODO(), request)
	if err != nil {
		return fmt.Errorf("starting stress test failed: %w", err)
	}

	log.Infof("Received: %v", reply)
	log.Info("✅ Starting stress test succeeded")

	return
}
