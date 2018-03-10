build:
	go build
dkbuild: build
	docker build -t reg.qiniu.com/wolfogre/qiniuauth:${version} .
dkpush:
	docker push reg.qiniu.com/wolfogre/qiniuauth:${version}
clean:
	rm -f qiniuauth
