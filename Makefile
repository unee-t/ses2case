all:
	@echo make {dev,demo,prod} to deploy

dev:
	@echo $$AWS_ACCESS_KEY_ID
	apex -r ap-southeast-1 --env dev deploy

demo:
	@echo $$AWS_ACCESS_KEY_ID
	apex -r ap-southeast-1 --env demo deploy

prod:
	@echo $$AWS_ACCESS_KEY_ID
	apex -r ap-southeast-1 --env prod deploy

logsdev:
	apex -r ap-southeast-1 --env dev logs

testdev:
	apex -r ap-southeast-1 --env dev invoke post < event.json

testdemo:
	apex -r ap-southeast-1 --env demo invoke post < event.json

testprod:
	apex -r ap-southeast-1 --env prod invoke post < event.json


.PHONY: dev demo prod testdev testdemo testprod
