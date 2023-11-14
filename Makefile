run:
	@go run . $(filter-out $@,$(MAKECMDGOALS))

%:
	@:

.PHONEY:
	run
