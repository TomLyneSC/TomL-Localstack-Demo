This is a demo showing off how we use localstack to make integration tests much easier to run, write and maintain

Steps to run:
- Make sure you have Localstack installed https://app.localstack.cloud/download

- Get all go dependencies
`go get -d ./...`

- Spin up localstack
`./scripts/setup-stack.sh`

You can now run the tests within the test suite!