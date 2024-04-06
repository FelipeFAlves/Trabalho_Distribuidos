genproto:
	protoc --proto_path=protob --go_out=contact --go_opt=paths=source_relative protob/*.proto