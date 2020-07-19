## Regenerate Protobufs: 

From this directory in a terminal, run:
`protoc --go_out=plugins=grpc:. *.proto`

We currently keep proto files in the separate repository, but I wonder if it's not better to just keep them in the service repo, like in this example. Are there any advantages to keeping them together?
