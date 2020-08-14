
run-eureka:
	docker run -p 8080:8080 --name eureka -d netflixoss/eureka:1.3.1

stop-eureka:
	docker stop eureka
	docker rm eureka

run-fulfill:
	go run cmd/fulfillment/main.go

run-catalog:
	go run cmd/catalog/main.go

test-catalog:
	curl localhost:8081/catalog/test