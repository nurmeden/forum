build:
	docker image build -f Dockerfile -t forum .
run:
	docker container run -p 8080:8080 --detach --name forum
allstop:
	docker stop $(docker ps -a -q)
prune:
	docker system prune -a