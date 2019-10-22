Скомпилировать go grpc: protoc protobuf/services.proto --go_out=plugins=grpc:./


Скомпилировать js grpc: protoc ./protobuf/services.proto --js_out=import_style=commonjs:./front-app/src/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./front-app/src/


$ npm install

$ npx webpack client.js

/* eslint-disable */