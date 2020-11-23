docker: job1 job2
	docker build . -f Dockerfile.job1 -t k8s4se/job1:latest --no-cache
	docker build . -f Dockerfile.job2 -t k8s4se/job2:latest --no-cache

job1: job1.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' job1.go

job2: job2.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' job2.go

clean:
	rm -f job1 job2 jobs.txt

push: docker
	docker push k8s4se/job1:latest
	docker push k8s4se/job2:latest
