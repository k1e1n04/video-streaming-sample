"use server";
import { video } from "@/proto/video";
import GetVideoRequest = video.GetVideoRequest;
import videoClient from "@/lib/grpcClient";
import UploadVideoRequest = video.UploadVideoRequest;

/**
 * Get video URL
 * @param videoId
 */
export const getVideoURL = async (videoId: string) => {
  return new Promise<string>((resolve, reject) => {
    const request = new GetVideoRequest();
    request.video_id = videoId;

    videoClient.GetVideoURL(request, {}, (err, response) => {
      if (err) {
        reject(err);
        return;
      }
      if (!response) {
        reject(new Error("No response"));
        return;
      }
      resolve(response.presigned_url);
    });
  });
};

/**
 * Upload video
 * @param file
 * @param title
 */
export const uploadVideo = async (
  fileBuffer: ArrayBuffer,
  title: string,
): Promise<string> => {
  const file = Buffer.from(fileBuffer);
  return new Promise((resolve, reject) => {
    const stream = videoClient.UploadVideo((err, response) => {
      if (err) {
        reject(err);
        return;
      }
      if (!response) {
        reject(new Error("No response"));
        return;
      }
      resolve(response.video_id);
    });

    // 1. send title
    const titleRequest = new UploadVideoRequest();
    titleRequest.title = title;
    stream.write(titleRequest);

    // 2. send file by chunk
    const chunkSize = 1024 * 64; // 64KB
    for (let i = 0; i < file.length; i += chunkSize) {
      const chunkRequest = new UploadVideoRequest();
      chunkRequest.chunk = file.subarray(i, i + chunkSize);
      stream.write(chunkRequest);
    }

    // 3. end stream
    stream.end();
  });
};
