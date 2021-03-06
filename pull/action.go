package pull

import (
	"context"

	"github.com/swisscom/korp/docker_utils"
	korpio "github.com/swisscom/korp/io"

	"github.com/docker/docker/api/types"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	kust "sigs.k8s.io/kustomize/pkg/image"
)

// Action - struct for pull action
type Action struct {
	Io korpio.PullPushIo
}

// Pull - Pull Docker images listed in the kustomization file from remote to the local Docker registry
func (p *Action) Pull(c *cli.Context) error {

	kstPath := c.String("kustomization-path")
	username := c.String("username")
	password := c.String("password")

	dockerImages, loadErr := p.Io.LoadKustomizationFile(kstPath)
	if loadErr != nil {
		log.Error(loadErr)
		return loadErr
	}

	pullErr := p.pullDockerImages(dockerImages, username, password)
	if pullErr != nil {
		log.Error(pullErr)
		return pullErr
	}

	return nil
}

// pullDockerImages - Pull all Docker images from given list
func (p *Action) pullDockerImages(dockerImages []kust.Image, username, password string) error {

	if len(dockerImages) > 0 {

		ctx := context.Background()

		cli, cliErr := p.Io.OpenDockerClient()
		if cliErr != nil {
			// log.Error(cliErr)
			return cliErr
		}
		defer cli.Close()

		daemonErr := docker_utils.CheckDockerDaemon(cli, &ctx)
		if daemonErr != nil {
			// log.Error(daemonErr)
			return daemonErr
		}

		pullOk, pullKo := 0, 0
		for _, img := range dockerImages {
			if p.pullDockerImage(cli, &ctx, img.Name, img.NewTag, username, password) {
				pullOk++
			} else {
				pullKo++
			}
		}
		log.Infof("Total Docker images pulled: %d - Total Docker images pulls failed: %d", pullOk, pullKo)
	} else {
		log.Warn("No Docker images to pull")
	}

	return nil
}

// pullDockerImage -
func (p *Action) pullDockerImage(cli docker_utils.DockerClient, ctx *context.Context, imageName, imageTag, username, password string) bool {

	imageRef := docker_utils.BuildCompleteDockerImage(imageName, imageTag)
	pullOpts := &types.ImagePullOptions{}
	if len(username) > 0 {
		pullOpts.RegistryAuth = docker_utils.EncodeRegistryAuth(username, password)
	}
	pullErr := docker_utils.PullDockerImage(cli, ctx, imageName, imageTag, pullOpts, true)
	if pullErr != nil {
		log.Errorf("Error pulling Docker image %s: %s", imageRef, pullErr.Error())
		return false
	} else {
		log.Infof("%s image pulled", imageRef)
		return true
	}
}
