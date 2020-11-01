
Sample code taken from https://gowebexamples.com/hello-world/

Compile with:

    CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

Build then run the docker image with:

    docker build . -t k8s4se/hello
    docker run -it -p 8080:8080 k8s4se/hello

Then vist http://localhost:8080/toto, you should see `Hello, you've requested: /toto`.

Note the size of the image!

    λ> docker image ls
    REPOSITORY                                                 TAG                      IMAGE ID            CREATED             SIZE
    k8s4se/hello                                               latest                   60ef39dbdd1a        2 minutes ago       10.5MB


