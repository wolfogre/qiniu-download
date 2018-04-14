version=`git tag -l | tail -n 1`

build:
	go build
dkbuild: build
	docker build -t reg.qiniu.com/wolfogre/qiniu-download:${version} .
dkpush:
	docker push reg.qiniu.com/wolfogre/qiniu-download:${version}
clean:
	rm -f qiniuauth
