integration-tests:
	docker compose -f ./deployments/docker/docker-compose.test.yaml up integration_test --abort-on-container-exit --exit-code-from integration_test

unit-tests:
	docker compose -f ./deployments/docker/docker-compose.test.yaml up unit_test --abort-on-container-exit --exit-code-from unit_test

benchmarks:
	docker compose -f ./deployments/docker/docker-compose.test.yaml up benchmarks --abort-on-container-exit --exit-code-from benchmarks