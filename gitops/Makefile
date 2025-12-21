k-start:
	minikube start
	minikube addons enable ingress

k-stop:
	minikube stop
	minikube delete --all

f-start:
	flux install

f-stop:
	flux uninstall -s

k-apply:
	kubectl apply -k ./clusters/main

k-tunnel:
	kubectl port-forward service/traefik 8080:web

k-delete:
	kubectl delete -k ./clusters/main
