import grpc from "@grpc/grpc-js";
import { video } from "@/proto/video";
import VideoServiceClient = video.VideoServiceClient;

const videoClient = new VideoServiceClient(
  "localhost:50052",
  grpc.credentials.createInsecure(),
);

export default videoClient;
