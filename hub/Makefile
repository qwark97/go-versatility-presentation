PACKAGES=`go list ./... | egrep -v 'mocks'`

test:
	go test ${PACKAGES} -cover -coverprofile=c.out -coverpkg=./... && \
	egrep -v 'mocks' c.out > filtered.out && \
	go tool cover -html=filtered.out -o coverage.html && \
	rm c.out filtered.out