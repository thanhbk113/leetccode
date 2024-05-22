bench-test:
	go test -bench=. -benchmem -benchtime=10s ./test/leet_test.go -run=^$ -v