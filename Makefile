gen-proto:
	protoc --gogoslick_out=plugins=grpc:pkg/ -I=. api/gophkeeper/*.proto