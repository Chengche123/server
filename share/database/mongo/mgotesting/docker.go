package mgotesting

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const (
	image = "mongo"
	port  = "27017/tcp"
)

// RunWithMongoInDocker create a temp container for mongo test
func RunWithMongoInDocker(m *testing.M, mongoURL *string) int {
	ctx := context.Background()

	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}

	res, err := c.ContainerCreate(ctx,
		&container.Config{
			Image: image,
			ExposedPorts: nat.PortSet{
				port: {},
			},
		}, &container.HostConfig{
			PortBindings: nat.PortMap{
				nat.Port(port): []nat.PortBinding{
					{
						HostIP:   "127.0.0.1",
						HostPort: "0",
					},
				},
			},
		}, nil, nil, "")

	if err != nil {
		panic(err)
	}

	defer func() {
		err = c.ContainerRemove(ctx, res.ID, types.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			panic(err)
		}
	}()

	err = c.ContainerStart(ctx, res.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	insRes, err := c.ContainerInspect(ctx, res.ID)
	if err != nil {
		panic(err)
	}

	addr := insRes.NetworkSettings.Ports[port][0]
	*mongoURL = fmt.Sprintf("mongodb://%s:%s", addr.HostIP, addr.HostPort)

	return m.Run()
}
