const redis = require("redis");
const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");

const args = {
  port: 8060,
  max_cache: 64,
};

process.argv.slice(2).map((arg) => {
  args[arg.split("=")[0]] = arg.split("=")[1];
});

const redis_client = redis.createClient({
  host: "redis", // For running with docker
  // host: 'localhost', // For running individually
  port: 6379,
});

redis_client.on("error", (err) => {
  console.log("Redis Error " + err);
});

const packageDefinition = protoLoader.loadSync("notes.proto");
const notesProto = grpc.loadPackageDefinition(packageDefinition);

const server = new grpc.Server();

server.bind(`127.0.0.1:${args.port}`, grpc.ServerCredentials.createInsecure());

console.log(`gprc server is running at 127.0.0.1:${args.port}`);

let notes = {};

server.addService(notesProto.CacheService.service, {
  getKey: (call, callback) => {
    const key = call.request.val;
    callback(null, notes);
  },

  setKey: (call, callback) => {
    const note = call.request;
    callback(null, note);
  },

  clear: (call, callback) => {
    callback(null, null);
  },
});

server.start();
