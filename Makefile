run:
	@go run . $(filter-out $@,$(MAKECMDGOALS))
install:
	go install .

%:
	@:

.PHONEY:
	run
