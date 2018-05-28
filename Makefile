build:
	docker image rm --force passion
	docker build -t passion .

run:
	docker run --rm -t -e SLACK_PASSION_KEY=$(SLACK_PASSION_KEY) -e GOOGLE_PASSION_KEY=$(GOOGLE_PASSION_KEY) -e CSE_ID=$(CSE_ID) passion

deploy:
	# arukas
	docker image rm --force tomokik/passion
	docker build --tag=tomokik/passion .
	docker push tomokik/passion

	# aws
	# docker image rm --force passion-bot
	# docker build -t passion-bot .
	# docker tag passion-bot:latest 861612634917.dkr.ecr.us-east-2.amazonaws.com/passion-bot:latest
	# docker push 861612634917.dkr.ecr.us-east-2.amazonaws.com/passion-bot:latest
