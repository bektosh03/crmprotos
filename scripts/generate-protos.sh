#!/bin/zsh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    teacherpb/teacher.proto studentpb/student.proto \
    schedulepb/schedule.proto journalpb/journal.proto