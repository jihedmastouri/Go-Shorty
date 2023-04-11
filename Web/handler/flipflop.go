package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/FlipFlop/proto"
	"google.golang.org/grpc"
)


func CallService(c echo.Context, conn *grpc.ClientConn) error {
	client := pb.NewCommandsServiceClient(conn)
	req := &pb.CreateCommandRequest{Name: "Mj", Description: "LoL"}
	c.Echo().Logger.Info("Request Processing")
	fmt.Println("request processing")
	res, err := client.CreateCommand(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res.GetId())
	return c.String(http.StatusOK, res.GetId())
}
