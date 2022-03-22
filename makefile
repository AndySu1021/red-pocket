server:
	go run main.go server

gen.mock:
	mockgen -source util/interface/http.go -destination util/mock/http_mock.go -package mock
	mockgen -source util/interface/cache.go -destination util/mock/cache_mock.go -package mock
