import { join } from "path";

let PROTO_PATH = "../../../wizd/src/api/packages.proto";

PROTO_PATH = join(__dirname, PROTO_PATH);

const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
});
const definitions = grpc.loadPackageDefinition(packageDefinition);
export const client = new definitions.commands.Packages(
  "localhost:50051",
  grpc.credentials.createInsecure()
);
