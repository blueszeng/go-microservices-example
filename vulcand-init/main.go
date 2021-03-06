package main

import "github.com/agencyrevolution/go-microservices-example/utils"

func main() {
	// Upsert backends
	utils.VulcandUpsertBackend("service.user")
	utils.VulcandUpsertBackend("service.repo")

	// Upsert frontends
	utils.VulcandUpsertFrontend("getUsers", "/users", "service.user")
	utils.VulcandUpsertFrontend("getUser", "/users/<username>", "service.user")
	utils.VulcandUpsertFrontend("getReposByUsername", "/users/<username>/repos", "service.repo")
	utils.VulcandUpsertFrontend("getRepo", "/users/<username>/repos/<reponame>", "service.repo")

	// Upsert listener
	utils.VulcandUpsertListener("demo", "", "0.0.0.0:3004")
}
