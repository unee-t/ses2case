REGION:=us-west-2

all:
	@echo make {dev,demo,prod} to deploy

dev:
	@echo $$AWS_ACCESS_KEY_ID
	apex -r $(REGION) --env dev deploy

demo:
	@echo $$AWS_ACCESS_KEY_ID
	apex -r $(REGION) --env demo deploy

prod:
	@echo $$AWS_ACCESS_KEY_ID
	apex -r $(REGION) --env prod deploy

logsdev:
	apex -r $(REGION) --env dev logs

testdev:
	apex -r $(REGION) --env dev invoke post < event.json

testdemo:
	apex -r $(REGION) --env demo invoke post < event.json

testprod:
	apex -r $(REGION) --env prod invoke post < event.json

testlocal:
	curl -i -H "Content-Type: application/json" -H "Authorization: Bearer blablabla" -X POST -d @event.json http://localhost:3000/api/ses

.PHONY: dev demo prod testdev testdemo testprod
