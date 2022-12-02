run:
	garble -debugdir ./garble build obfuscate/sha256
	var=$(<./sha256/sha256.go) && mumbojumbo -s="$var" -p="sha256" | goimports > ./mumbojumbo/sha256.go
	go run main.go
	poltergeist -encode -file ./sha256/sha256.go > polter.go