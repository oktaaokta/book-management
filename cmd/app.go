package main

func main() {
	repositories := initRepository()
	usecases := initUsecase(repositories)
	handler := initHandler(usecases)

	serveHTTP(handler)
}
