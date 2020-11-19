docker: crashloop_reporter
	docker build . -t k8s4se/crashloop-reporter:latest --no-cache

crashloop_reporter: crashloop_reporter.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

clean:
	rm -f crashloop_reporter

push: docker
	docker push k8s4se/crashloop-reporter
