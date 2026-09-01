package deploy

const DistrolessDockerfile = `FROM gcr.io/distroless/static-debian12:nonroot
COPY gleedos /usr/local/bin/gleedos
ENTRYPOINT ["/usr/local/bin/gleedos"]
`

func GetDistrolessDockerfile() string {
	return DistrolessDockerfile
}
