# CPU Load Generator

A Program to generate CPU Load on Unices (Tried on mac OSX) and Windows, trying to attain a Threshold (30.0 % default)

clone the repository, then:
```shell
cd cpu-load-generator
go get .
go build && go install
cpu-load-generator --help 
```

Code for CPU Time accounting is in cpu_tools_{darwin,windows}.
