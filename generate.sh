protoc --go_out=src/ src/simple.proto
protoc --go_out=greet/greetpb/ greet/greetpb/greet.proto \
--go-grpc_out=greet/greetpb/ greet/greetpb/greet.proto