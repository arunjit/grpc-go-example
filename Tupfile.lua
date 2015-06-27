tup.export('GOROOT')
tup.export('GOPATH')

tup.definerule{
  inputs = {'proto/greet.proto'},
  command = 'protoc --go_out=plugins=grpc:. proto/greet.proto',
  outputs = {'proto/greet.pb.go'},
}

tup.definerule{
  inputs = {'server/main.go', 'proto/greet.pb.go'},
  command = 'go build -o build/server server/main.go',
  outputs = {'build/server'},
}

tup.definerule{
  inputs = {'client/main.go', 'proto/greet.pb.go'},
  command = 'go build -o build/client client/main.go',
  outputs = {'build/client'},
}
