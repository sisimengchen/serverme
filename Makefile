default: build

build:
	go fmt
	cd ./app 
	go fmt
	cd ..
	cd ./configs 
	go fmt
	cd ..
	cd ./controllers 
	go fmt
	cd ..
	cd ./database 
	go fmt
	cd ..
	cd ./middleware 
	go fmt
	cd ..
	cd ./models 
	go fmt
	cd ..
	cd ./routes 
	go fmt
	cd ..
	cd ./utils 
	go fmt
	cd ..
	