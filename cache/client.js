const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");

const PROTO_PATH = "./notes.proto";
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const TodoService = grpc.loadPackageDefinition(packageDefinition).TodoService;

const client = new TodoService(
  "localhost:50051",

  grpc.credentials.createInsecure()
);

module.exports = client;
