.PHONY: build

clean: clean_local_registry

clean_local_registry:
	docker system prune -f

build:
	docker build -t registry.cn-qingdao.aliyuncs.com/wqer/hello-k8s:latest .

upload: clean build
	docker push registry.cn-qingdao.aliyuncs.com/wqer/hello-k8s:latest