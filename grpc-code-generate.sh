for element in `ls -F api | grep "/$"`;
do
  protoc --go_out=plugins=grpc,paths=source_relative:. ./api/${element}/*.proto
done
