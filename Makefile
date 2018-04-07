build:
	docker image rm --force passion
	docker build -t passion .

run:
	docker run -t -e SLACK_PASSION_KEY=$(SLACK_PASSION_KEY) passion
