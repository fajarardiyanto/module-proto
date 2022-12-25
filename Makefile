# Build
go-proto:
	@bash -c "source ./scripts/before.sh && ./scripts	pts/go-proto.sh"
go-proto-gen-docker:
	@scripts/docker.sh