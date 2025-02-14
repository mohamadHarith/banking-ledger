 protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    shared/proto/transaction_processor_proto/transaction_processor_service.proto