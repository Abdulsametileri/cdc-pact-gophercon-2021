#make consumer VERSION="1.0.0"
consumer:
	rm -rf product/pacts
	go test ./... -tags=consumer -count=1
	sh product/pact-publish.sh $(VERSION)

provider:
	go test ./... -tags=provider