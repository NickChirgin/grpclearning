protoc --go_out=src/ src/simple.proto
protoc --go_out=greet/greetpb/ greet/greetpb/greet.proto \
--go-grpc_out=greet/greetpb/ greet/greetpb/greet.proto
protoc --go_out=calculator calculator/calculatorpb/calculator.proto \
--go-grpc_out=calculator calculator/calculatorpb/calculator.proto