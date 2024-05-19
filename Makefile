build:
	go build -o go-cron-parse .
run:
	go run .

# Examples
ex1:
	go run . "* * * * * program"
ex2:
	go run . "*/15 0 1,15 * 1-5 /usr/bin/find"
ex3:
	go run . "*/15 0 1,15 * mon-tuesday /usr/bin/find"
ex4:
	go run . "*/15 0 1,15 * mon-tuesday /usr/bin/find"
ex5:
	go run . "*/15 0 1,15 jan-mar mon-tuesday /usr/bin/find"
ex6:
	go run . "*/15 0 1-15/5 jan-mar mon-tuesday /usr/bin/find"
