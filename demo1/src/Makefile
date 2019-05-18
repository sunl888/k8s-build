.PHONY: build

clean: clean_docker

clean_docker:
	docker system prune -f

build_docker:
	docker build -t registry.cn-qingdao.aliyuncs.com/wqer/hello-k8s:latest .

upload: clean build_docker
	docker push registry.cn-qingdao.aliyuncs.com/wqer/hello-k8s:latest