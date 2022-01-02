const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");

const uuid = require("uuid");

const packageDefinition = protoLoader.loadSync("notes.proto");
const notesProto = grpc.loadPackageDefinition(packageDefinition);

const server = new grpc.Server();

server.bind("127.0.0.1:50051", grpc.ServerCredentials.createInsecure());

console.log("server is running at http://127.0.0.1:50051");

const todos = [];

server.addService(notesProto.TodoService.service, {
  list: (_, callback) => {
    callback(null, todos);
  },

  insert: (call, callback) => {
    let todo = call.request;

    todo.id = uuid.v4();

    todos.push(todo);

    callback(null, todo);
  },
});

server.start();
