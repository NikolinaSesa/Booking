module github.com/NikolinaSesa/Booking/api-gateway

go 1.20

replace (
	github.com/NikolinaSesa/Booking/apartment-service => ../apartment-service
	github.com/NikolinaSesa/Booking/user-service => ../user-service
)

require (
	github.com/NikolinaSesa/Booking/user-service v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0
	google.golang.org/grpc v1.56.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
