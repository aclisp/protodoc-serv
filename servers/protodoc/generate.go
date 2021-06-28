package protodoc

//go:generate protoc --micro_out=. --go_out=. protodoc.proto
//go:generate bash -c "protodoc protodoc.proto > protodoc.md"
