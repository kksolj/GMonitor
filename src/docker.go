package main

import (
	"context"
	"github.com/docker/docker/api/types"
	client2 "github.com/docker/docker/client"
)

/*func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		log.Printf("%s %s %+v\n", container.ID[0:10], container.Image, container.Names)
	}
}*/
var client *client2.Client

func InitClient() {
	var err error
	client, err = client2.NewEnvClient()
	if err != nil {
		panic(err)
	}
}
func tryInitClient() {
	defer func() {
		if e := recover(); e != nil {
			return
		}
	}()
	InitClient()
}
func containers(all bool) (c []types.Container) {
	c, _ = client.ContainerList(context.Background(), types.ContainerListOptions{
		All: all,
	})
	return
}
func startContainer(id string) {
	for _, c := range containers(true) {
		if c.ID == id {
			client.ContainerStart(context.Background(), c.ID, types.ContainerStartOptions{})
		}
	}
}
func startContainerByName(name string) {
	for _, c := range containers(true) {
		if sliceContains(c.Names, name) {
			client.ContainerStart(context.Background(), c.ID, types.ContainerStartOptions{})
		}
	}
}

func sliceContains(str []string, s string) bool {
	for _, x := range str {
		if x == s {
			return true
		}
	}
	return false
}
