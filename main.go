package styx

import (
	"context"

	"fmt"

	"github.com/solher/styx/pb"
	"google.golang.org/grpc"
)

// CreateSession is ...
func CreateSession(authServerURL string, payload []byte, policies []string) (*pb.Session, error) {
	conn, err := grpc.Dial(authServerURL, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		return nil, err
	}
	c := pb.NewSessionManagementClient(conn)
	ctx := context.Background()

	session := &pb.Session{
		Policies: policies,
		Payload:  payload,
	}

	reply, err := c.CreateSession(ctx, &pb.CreateSessionRequest{Session: session})
	if err != nil {
		return nil, err
	}

	return reply.GetSession(), nil
}

// CreateSessionWithToken is ...
func CreateSessionWithToken(authServerURL string, token string, payload []byte, policies []string) (*pb.Session, error) {
	conn, err := grpc.Dial(authServerURL, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	client := pb.NewSessionManagementClient(conn)
	ctx := context.Background()

	session := &pb.Session{
		Policies:   policies,
		Payload:    payload,
		OwnerToken: token,
	}

	fmt.Println(session)

	reply, err := client.CreateSession(ctx, &pb.CreateSessionRequest{Session: session})
	if err != nil {
		return nil, err
	}

	return reply.GetSession(), nil
}

// GetSession is ...
func GetSession(authServerURL string, token string) (*pb.Session, error) {
	conn, err := grpc.Dial(authServerURL, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	client := pb.NewSessionManagementClient(conn)
	ctx := context.Background()

	reply, err := client.FindSessionByToken(ctx, &pb.FindSessionByTokenRequest{Token: token})
	if err != nil {
		return nil, err
	}

	return reply.GetSession(), nil
}

// DeleteSession is ...
func DeleteSession(authServerURL string, token string) (*pb.Session, error) {
	conn, err := grpc.Dial(authServerURL, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	client := pb.NewSessionManagementClient(conn)
	ctx := context.Background()

	reply, err := client.DeleteSessionByToken(ctx, &pb.DeleteSessionByTokenRequest{Token: token})
	if err != nil {
		return nil, err
	}

	return reply.GetSession(), nil

}
