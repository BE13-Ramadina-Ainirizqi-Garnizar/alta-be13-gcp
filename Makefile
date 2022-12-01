test:
	go test ./feature/user/... -coverprofile=cover.out && go tool cover -html=cover.out