// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/swisscom/korp/docker_utils"
	"io"
	"sync"
)

var (
	lockDockerClientMockClose         sync.RWMutex
	lockDockerClientMockContainerList sync.RWMutex
	lockDockerClientMockImagePull     sync.RWMutex
	lockDockerClientMockImagePush     sync.RWMutex
	lockDockerClientMockImageTag      sync.RWMutex
)

// Ensure, that DockerClientMock does implement DockerClient.
// If this is not the case, regenerate this file with moq.
var _ docker_utils.DockerClient = &DockerClientMock{}

// DockerClientMock is a mock implementation of DockerClient.
//
//     func TestSomethingThatUsesDockerClient(t *testing.T) {
//
//         // make and configure a mocked DockerClient
//         mockedDockerClient := &DockerClientMock{
//             CloseFunc: func() error {
// 	               panic("mock out the Close method")
//             },
//             ContainerListFunc: func(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
// 	               panic("mock out the ContainerList method")
//             },
//             ImagePullFunc: func(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {
// 	               panic("mock out the ImagePull method")
//             },
//             ImagePushFunc: func(ctx context.Context, image string, options types.ImagePushOptions) (io.ReadCloser, error) {
// 	               panic("mock out the ImagePush method")
//             },
//             ImageTagFunc: func(ctx context.Context, source string, target string) error {
// 	               panic("mock out the ImageTag method")
//             },
//         }
//
//         // use mockedDockerClient in code that requires DockerClient
//         // and then make assertions.
//
//     }
type DockerClientMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// ContainerListFunc mocks the ContainerList method.
	ContainerListFunc func(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)

	// ImagePullFunc mocks the ImagePull method.
	ImagePullFunc func(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)

	// ImagePushFunc mocks the ImagePush method.
	ImagePushFunc func(ctx context.Context, image string, options types.ImagePushOptions) (io.ReadCloser, error)

	// ImageTagFunc mocks the ImageTag method.
	ImageTagFunc func(ctx context.Context, source string, target string) error

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// ContainerList holds details about calls to the ContainerList method.
		ContainerList []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Options is the options argument value.
			Options types.ContainerListOptions
		}
		// ImagePull holds details about calls to the ImagePull method.
		ImagePull []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// RefStr is the refStr argument value.
			RefStr string
			// Options is the options argument value.
			Options types.ImagePullOptions
		}
		// ImagePush holds details about calls to the ImagePush method.
		ImagePush []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Image is the image argument value.
			Image string
			// Options is the options argument value.
			Options types.ImagePushOptions
		}
		// ImageTag holds details about calls to the ImageTag method.
		ImageTag []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Source is the source argument value.
			Source string
			// Target is the target argument value.
			Target string
		}
	}
}

// Close calls CloseFunc.
func (mock *DockerClientMock) Close() error {
	if mock.CloseFunc == nil {
		panic("DockerClientMock.CloseFunc: method is nil but DockerClient.Close was just called")
	}
	callInfo := struct {
	}{}
	lockDockerClientMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockDockerClientMockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedDockerClient.CloseCalls())
func (mock *DockerClientMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockDockerClientMockClose.RLock()
	calls = mock.calls.Close
	lockDockerClientMockClose.RUnlock()
	return calls
}

// ContainerList calls ContainerListFunc.
func (mock *DockerClientMock) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	if mock.ContainerListFunc == nil {
		panic("DockerClientMock.ContainerListFunc: method is nil but DockerClient.ContainerList was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Options types.ContainerListOptions
	}{
		Ctx:     ctx,
		Options: options,
	}
	lockDockerClientMockContainerList.Lock()
	mock.calls.ContainerList = append(mock.calls.ContainerList, callInfo)
	lockDockerClientMockContainerList.Unlock()
	return mock.ContainerListFunc(ctx, options)
}

// ContainerListCalls gets all the calls that were made to ContainerList.
// Check the length with:
//     len(mockedDockerClient.ContainerListCalls())
func (mock *DockerClientMock) ContainerListCalls() []struct {
	Ctx     context.Context
	Options types.ContainerListOptions
} {
	var calls []struct {
		Ctx     context.Context
		Options types.ContainerListOptions
	}
	lockDockerClientMockContainerList.RLock()
	calls = mock.calls.ContainerList
	lockDockerClientMockContainerList.RUnlock()
	return calls
}

// ImagePull calls ImagePullFunc.
func (mock *DockerClientMock) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {
	if mock.ImagePullFunc == nil {
		panic("DockerClientMock.ImagePullFunc: method is nil but DockerClient.ImagePull was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		RefStr  string
		Options types.ImagePullOptions
	}{
		Ctx:     ctx,
		RefStr:  refStr,
		Options: options,
	}
	lockDockerClientMockImagePull.Lock()
	mock.calls.ImagePull = append(mock.calls.ImagePull, callInfo)
	lockDockerClientMockImagePull.Unlock()
	return mock.ImagePullFunc(ctx, refStr, options)
}

// ImagePullCalls gets all the calls that were made to ImagePull.
// Check the length with:
//     len(mockedDockerClient.ImagePullCalls())
func (mock *DockerClientMock) ImagePullCalls() []struct {
	Ctx     context.Context
	RefStr  string
	Options types.ImagePullOptions
} {
	var calls []struct {
		Ctx     context.Context
		RefStr  string
		Options types.ImagePullOptions
	}
	lockDockerClientMockImagePull.RLock()
	calls = mock.calls.ImagePull
	lockDockerClientMockImagePull.RUnlock()
	return calls
}

// ImagePush calls ImagePushFunc.
func (mock *DockerClientMock) ImagePush(ctx context.Context, image string, options types.ImagePushOptions) (io.ReadCloser, error) {
	if mock.ImagePushFunc == nil {
		panic("DockerClientMock.ImagePushFunc: method is nil but DockerClient.ImagePush was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Image   string
		Options types.ImagePushOptions
	}{
		Ctx:     ctx,
		Image:   image,
		Options: options,
	}
	lockDockerClientMockImagePush.Lock()
	mock.calls.ImagePush = append(mock.calls.ImagePush, callInfo)
	lockDockerClientMockImagePush.Unlock()
	return mock.ImagePushFunc(ctx, image, options)
}

// ImagePushCalls gets all the calls that were made to ImagePush.
// Check the length with:
//     len(mockedDockerClient.ImagePushCalls())
func (mock *DockerClientMock) ImagePushCalls() []struct {
	Ctx     context.Context
	Image   string
	Options types.ImagePushOptions
} {
	var calls []struct {
		Ctx     context.Context
		Image   string
		Options types.ImagePushOptions
	}
	lockDockerClientMockImagePush.RLock()
	calls = mock.calls.ImagePush
	lockDockerClientMockImagePush.RUnlock()
	return calls
}

// ImageTag calls ImageTagFunc.
func (mock *DockerClientMock) ImageTag(ctx context.Context, source string, target string) error {
	if mock.ImageTagFunc == nil {
		panic("DockerClientMock.ImageTagFunc: method is nil but DockerClient.ImageTag was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Source string
		Target string
	}{
		Ctx:    ctx,
		Source: source,
		Target: target,
	}
	lockDockerClientMockImageTag.Lock()
	mock.calls.ImageTag = append(mock.calls.ImageTag, callInfo)
	lockDockerClientMockImageTag.Unlock()
	return mock.ImageTagFunc(ctx, source, target)
}

// ImageTagCalls gets all the calls that were made to ImageTag.
// Check the length with:
//     len(mockedDockerClient.ImageTagCalls())
func (mock *DockerClientMock) ImageTagCalls() []struct {
	Ctx    context.Context
	Source string
	Target string
} {
	var calls []struct {
		Ctx    context.Context
		Source string
		Target string
	}
	lockDockerClientMockImageTag.RLock()
	calls = mock.calls.ImageTag
	lockDockerClientMockImageTag.RUnlock()
	return calls
}
