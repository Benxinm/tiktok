DIR = $(shell pwd)
CMD = $(DIR)/cmd
OUTPUT_PATH = $(DIR)/output

SERVICES := api user follow interaction video chat
service = $(word 1, $@)

env-up:
	docker-compose up -d

env-down:
	docker-compose down

$(SERVICES):
	mkdir -p output
	cd $(CMD)/$(service) && sh build.sh
	cd $(CMD)/$(service)/output && mv . $(OUTPUT_PATH)/$(service)
ifdef ci
	sh $(OUTPUT_PATH)/$(service)/bootstrap.sh $(ENPOINT)
endif

build-all:
	@for var in $(SERVICES); do \
  		make $$var ci=1; \
  	done