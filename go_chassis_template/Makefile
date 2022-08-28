proto:
	set -ue
	cd protocol/grpc/ && mkdir -p go && protoc -I ./ --go_out=plugins=grpc:./go/ *.proto

lint:
	if command -v golangci-lint >/dev/null 2>&1;then \
		echo "exists golangci-lint"; \
	else \
		echo "not exists golangci-lint"; \
		brew install golangci-lint; \
	fi
	golangci-lint run -v    --skip-dirs="vendor" --skip-files "_test\.go" --exclude unused,sa1012,asmdecl,S1004,SA5008,SA1029  ./...

pb:
	docker run --rm -it  -v $(PWD):$(PWD) -w $(PWD) go_scaffold:v1 make proto

run_grpc:
	CHASSIS_HOME=$(PWD) CHASSIS_CONF_DIR=$(PWD)/conf/grpc go run cmd/grpc/main.go