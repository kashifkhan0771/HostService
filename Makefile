format:
	bash ./scripts/format.sh
check: format
	bash ./scripts/check.sh
test: check
	bash ./scripts/test.sh