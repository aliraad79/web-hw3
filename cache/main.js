import { createClient } from "redis";
import grpc from "grpc";
import protoLoader from "@grpc/proto-loader";

const args = {
  port: 8060,
  max_cache: 64,
};

process.argv.slice(2).map((arg) => {
  args[arg.split("=")[0]] = arg.split("=")[1];
});

const redisClient = createClient({
  // host: "redis", // For running with docker
  host: "localhost", // For running individually
  port: 6379,
});

redisClient.on("error", (err) => console.log("Redis Client Error", err));

await redisClient.connect();

const packageDefinition = protoLoader.loadSync("notes.proto");
const notesProto = grpc.loadPackageDefinition(packageDefinition);

const server = new grpc.Server();

server.bind(`127.0.0.1:${args.port}`, grpc.ServerCredentials.createInsecure());

console.log(`gprc server is running at 127.0.0.1:${args.port}`);

server.addService(notesProto.CacheService.service, {
  getKey: async (call, callback) => {
    const key = call.request.val;
    if (!(await redisClient.exists(key))) {
      return callback({
        code: 400,
        message: "miss cache",
        status: grpc.status.UNAVAILABLE,
      });
    }
    const value = await redisClient.get(key);
    return callback(null, { val: value });
  },

  setKey: async (call, callback) => {
    const note = call.request;
    await redisClient.set(note.key, note.value);
    callback(null, note);
  },

  clear: (call, callback) => {
    redisClient.flushAll();
    callback(null, null);
  },
});

server.start();
