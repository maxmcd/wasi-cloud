


python-0.1.0.tar.gz:
	wget https://registry-cdn.wapm.io/packages/_/python/python-0.1.0.tar.gz

python: python-0.1.0.tar.gz
	mkdir python
	cd python && tar fxz ../python-0.1.0.tar.gz

.PHONY: run
run: python
	go run main.go
