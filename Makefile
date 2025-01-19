ROOT_DIR    = $(shell pwd)
TEST_CONFIG_FILE = "config.test.yaml"
ADMIN_RESOURCE_PATH = "/resource/public/admin/"

########################
# Tools
########################
.PHONY: tools tools.gotools tools.gf tools.buf tools.ansible
# install dev tools
tools: tools.gotools

# install tools use go install
# go install have caches, so we don't need to check if it's installed
tools.gotools:
	@echo "Install Go Tools"
	@go mod download
	@./scripts/tools/tools.sh
	@echo "Install Go Tools Done"

tools.buf:
	@set -e; \
	if ! buf --version > /dev/null 2>&1; then \
  		echo "Buf CLI is not installed, start proceeding auto installation..."; \
		make tools; \
	fi;

tools.atlas:
	@set -e; \
	if ! atlas version > /dev/null 2>&1; then \
  		echo "Atlas CLI is not installed, start proceeding auto installation..."; \
		curl -sSf https://atlasgo.sh | sh; \
	fi;


########################
# Dev
########################
.PHONY: dev dev.up

# setup dev environment for first time
# 1. install dev tools
# 2. setup db, redis, etc.
dev: tools dev.up

# setup dev environment for db, redis, etc.
dev.up:
	@echo "Start Dev Environment"
	@./scripts/dev/up.sh
	@echo "Start Dev Environment Done"

dev.down:
	@echo "Stop Dev Environment"
	@./scripts/dev/down.sh
	@echo "Stop Dev Environment Done"

########################
# Proto
########################
.PHONY: proto proto.up proto.gen
proto: proto.gen

# update proto dependencies
proto.up: tools.buf
	@echo "Update Proto Dependencies"
	@buf dep update
	@echo "Update Proto Dependencies Done"

# generate proto files
proto.gen: tools
	@echo "Generate Proto Files"
	@cd proto && buf generate
	@echo "Generate Proto Files Done"


########################
# Migrate
########################
.PHONY: migrate migrate.mysql migrate.postgres migrate.sqlite3
migrate: migrate.mysql migrate.postgres migrate.sqlite3

migrate.mysql: tools.atlas
	@echo "Apply Migrate"
	@MIGRATE_ENV=mysql \
	MIGRATE_MODE=auto-approve \
	./scripts/migrate/apply.sh
	@echo "Apply Migrate Done"

migrate.postgres: tools.atlas
	@echo "Apply Migrate"
	@MIGRATE_ENV=postgres \
	MIGRATE_MODE=auto-approve \
	./scripts/migrate/apply.sh
	@echo "Apply Migrate Done"

migrate.sqlite3: tools.atlas
	@echo "Apply Migrate"
	@MIGRATE_ENV=sqlite3 \
	MIGRATE_MODE=auto-approve \
	./scripts/migrate/apply.sh
	@echo "Apply Migrate Done"
