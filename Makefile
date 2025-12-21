GITHUB_USER ?= SadraSamadi

kube-up:
	minikube start

kube-down:
	minikube stop
	minikube delete --all

flux-boot:
	flux bootstrap github \
      --token-auth \
      --owner=$(GITHUB_USER) \
      --repository=cloud-play \
      --branch=main \
      --path=k8s \
      --personal
