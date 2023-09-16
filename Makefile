test:
	echo "run tests recording the HTTP interaction"
	go test ./user/replayer -tags golden -count=1 -v
	echo "run tests replaying the HTTP interaction"
	go test ./user/replayer -count=1 -v