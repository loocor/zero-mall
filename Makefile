RPC-User:
	docker exec docker-golang-1 go run service/user/rpc/user.go -f service/user/rpc/etc/user.yaml

API-User:
	docker exec docker-golang-1 go run service/user/api/user.go -f service/user/api/etc/user.yaml

RPC-Product:
	docker exec docker-golang-1 go run service/product/rpc/product.go -f service/product/rpc/etc/product.yaml

API-Product:
	docker exec docker-golang-1 go run service/product/api/product.go -f service/product/api/etc/product.yaml

RPC-Order:
	docker exec docker-golang-1 go run service/order/rpc/order.go -f service/order/rpc/etc/order.yaml

API-Order:
	docker exec docker-golang-1 go run service/order/api/order.go -f service/order/api/etc/order.yaml

RPC-Pay:
	docker exec docker-golang-1 go run service/pay/rpc/pay.go -f service/pay/rpc/etc/pay.yaml

API-Pay:
	docker exec docker-golang-1 go run service/pay/api/pay.go -f service/pay/api/etc/pay.yaml

# DockerRestart:
# 	docker container restart docker-golang-1

Prometheus&Jaeger:
	open http://localhost:3000/targets
	open http://localhost:5000/search