package main

import (
	"context"
)

type Fedora struct{
	// Returns the Fedora image repository
	Image string
}

func New(
	// The Fedora image for the container to run
	// +optional
	// +default="docker.io/fedora:40"
	image string,
) *Fedora {
	return &Fedora{
		Image: image,
	}
}

// Build a Fedora container
func (m *Fedora) Container(
	// DNF packages to install
	// +optional
	packages []string,
	// Overlay images to merge on top of the base
	// +optional
	overlays []*Container,
) *Container {
	installCmd := []string{
		"dnf",
		"--assumeyes",
		"install",
		"--setopt=install_weak_deps=0",
		"--setopt=tsflags=nodocs",
	}

	ctr := dag.Container().From(m.Image)
	if len(packages) > 0 {
		ctr = ctr.WithExec(append(installCmd, packages...))
	}

	for _, overlay := range overlays {
		ctr = ctr.WithDirectory("/", overlay.Rootfs())
	}
	return ctr
}

// Returns the operating system identification data from /etc/os-release
func (m *Fedora) OsRelease(ctx context.Context) (string, error) {
	return m.Container(make([]string,0), make([]*Container,0)).WithExec([]string{"cat", "/etc/os-release"}).Stdout(ctx)
}


