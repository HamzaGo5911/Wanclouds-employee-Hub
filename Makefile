MONGO_INIT_CONTAINER_NAME = mongo-init-container

mymongoinit:
	@if [ ! $$(docker ps -aq -f name=$(MONGO_INIT_CONTAINER_NAME)) ]; then \
		docker run --name $(MONGO_INIT_CONTAINER_NAME) -d -p 27017:27017 mongo; \
	fi
	sleep 5

db_stop:
	bash ./scripts/db_stop.sh

test:db_stop mymongoinit
	bash ./scripts/test.sh

.PHONY: mymongoinit db_stop test
