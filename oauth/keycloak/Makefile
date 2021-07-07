


# tag::run-keycloak[]
.PHONY: run-keycloak
run-keycloak:
	docker rm -f mykeycloak
	docker run -d -p 8080:8080 --name mykeycloak -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin jboss/keycloak
# end::run-keycloak[]