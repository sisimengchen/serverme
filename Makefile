default: build

build:
	go fmt
	cd ./app && go fmt
	cd ./configs && go fmt
	cd ./controllers && go fmt 
	cd ./middleware && go fmt 
	cd ./models && go fmt 
	cd ./routes && go fmt 
	cd ./utils && go fmt 
	