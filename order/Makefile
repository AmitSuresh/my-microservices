# Define paths and variables
TERRAFORM = terraform
TF_OUTPUT_FILE = terraform/outputs.json
ENV_FILE = .env
GO_APP = main.go
TERRAFORM_DIR = terraform

# AWS Region
AWS_REGION = ap-southeast-1

# Target to initialize Terraform
terraform-init:
	@echo "Initializing Terraform..."
	$(TERRAFORM) -chdir=$(TERRAFORM_DIR) init -backend-config="region=$(AWS_REGION)"

# Target to create and review Terraform execution plan
terraform-plan: terraform-init
	@echo "Creating Terraform execution plan..."
	$(TERRAFORM) -chdir=$(TERRAFORM_DIR) plan -out=tfplan

# Target to apply Terraform configuration and provision resources
terraform-apply: terraform-plan
	@echo "Applying Terraform configuration..."
	$(TERRAFORM) -chdir=$(TERRAFORM_DIR) apply -auto-approve tfplan

# Export Terraform outputs to .env file (root folder)
terraform-outputs:
	@echo "Exporting Terraform outputs to .env file..."
	$(TERRAFORM) -chdir=$(TERRAFORM_DIR) output -json > $(TF_OUTPUT_FILE)
	jq -r 'to_entries | .[] | "\(.key | ascii_upcase)=\(.value.value)"' $(TF_OUTPUT_FILE) >> $(ENV_FILE)

# Ensure .env file exists and append outputs from Terraform
check-env-file:
	@echo "Checking if .env file exists in root folder..."
	[ ! -f $(ENV_FILE) ] && touch $(ENV_FILE)

# Target to run the Go application with environment variables from .env
run-go-app: check-env-file terraform-apply terraform-outputs
	@echo "Running Go application..."
	GO_ENV=$(ENV_FILE) go run $(GO_APP)

# Clean up generated Terraform outputs file
tf-clean:
	@echo "Cleaning up..."
	rm -f $(TF_OUTPUT_FILE)

# Target to destroy Terraform resources (stop incurring AWS costs)
terraform-destroy:
	@echo "Destroying Terraform-managed resources..."
	$(TERRAFORM) -chdir=$(TERRAFORM_DIR) destroy -auto-approve

# Full deploy process: create plan, apply plan, export outputs, and run the Go application
deploy: terraform-plan terraform-apply terraform-outputs run-go-app
	@echo "Deployment complete!"

# Full clean and destroy: clean outputs and destroy resources
destroy: terraform-destroy clean
	@echo "Terraform resources destroyed and cleaned up."

.PHONY: build run stop clean prune

build:
	@echo "Building Docker images..."
	sudo docker-compose build

run:
	@echo "Starting Docker containers..."
	sudo docker-compose up -d

stop:
	@echo "Stopping running containers..."
	@if [ ! -z "$$(sudo docker ps -q)" ]; then \
		sudo docker stop $$(sudo docker ps -q); \
	else \
		echo "No running containers to stop."; \
	fi

clean:
	@echo "Stopping and removing all containers..."
	@if [ ! -z "$$(sudo docker ps -aq)" ]; then \
		sudo docker stop $$(sudo docker ps -aq); \
		sudo docker rm $$(sudo docker ps -aq); \
	else \
		echo "No containers to clean."; \
	fi

# Remove all unused networks
clean-networks:
	@echo "Removing unused Docker networks..."
	@if [ ! -z "$$(docker network ls -q)" ]; then \
		sudo docker network rm $$(docker network ls -q); \
	else \
		echo "No networks to remove."; \
	fi

# Remove all unused volumes
clean-volumes:
	@echo "Removing all Docker volumes..."
	@if [ ! -z "$$(docker volume ls -q)" ]; then \
		sudo docker volume rm $$(docker volume ls -q); \
	else \
		echo "No volumes to remove."; \
	fi

# Prune Docker system (removes unused images, containers, volumes, etc.)
prune:
	@echo "Pruning Docker system..."
	sudo docker system prune -a --volumes -f

COMPOSE_FILE ?= docker-compose.yaml
.PHONY: compose-build
compose-build:
	docker-compose -f $(COMPOSE_FILE) build

.PHONY: compose-up
compose-up:
	docker compose -f $(COMPOSE_FILE) up

.PHONY: compose-up-build
compose-up-build:
	docker compose -f $(COMPOSE_FILE) up --build

.PHONY: compose-down
compose-down:
	docker compose -f $(COMPOSE_FILE) down

.PHONY: docker-rm-all
docker-rm-all:
	@if [ "$(docker ps -aq)" ]; then \
		docker stop $(docker ps -aq); \
		docker rm $(docker ps -aq); \
	fi
	@if [ "$(docker images -q)" ]; then \
		docker rmi $(docker images -q); \
	fi
	@if [ "$(docker network ls -q)" ]; then \
		docker network rm $(docker network ls -q); \
	fi
	@if [ "$(docker volume ls -q)" ]; then \
		docker volume rm $(docker volume ls -q); \
	fi
	docker system prune -a --volumes --force

