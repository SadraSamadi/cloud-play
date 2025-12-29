GITHUB_USER ?= SadraSamadi

kube-up:
	minikube start
	minikube addons enable metrics-server

kube-down:
	minikube stop
	minikube delete --all

flux-boot:
	flux bootstrap github \
		--components-extra=image-reflector-controller,image-automation-controller \
		--token-auth \
		--owner=$(GITHUB_USER) \
		--repository=cloud-play \
		--branch=main \
		--path=k3s \
		--read-write-key \
		--personal
