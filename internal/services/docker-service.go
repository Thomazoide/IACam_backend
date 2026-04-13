package services

import (
	"context"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var cli *client.Client

func InitDocker() {
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
}

func CreateWorker(cameraID uint, rtsp string) error {
	containerName := "worker-camera-" + strconv.Itoa(int(cameraID))
	resp, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: "python-worker",
			Env: []string{
				"RTSP_URL=" + rtsp,
				"CAMERA_ID=" + strconv.Itoa(int(cameraID)),
				"API_URL=http://localhost:8080/events",
			},
		},
		nil, nil, nil,
		containerName,
	)
	if err != nil {
		return err
	}
	return cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
}

func RemoveWorker(cameraID uint) error {
	containerName := "worker-camera-" + strconv.Itoa(int(cameraID))
	timeout := 5
	if err := cli.ContainerStop(context.Background(), containerName, container.StopOptions{Timeout: &timeout}); err != nil {
		return err
	}
	return cli.ContainerRemove(context.Background(), containerName, types.ContainerRemoveOptions{})
}
