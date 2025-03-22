"use server";
import { video } from "@/proto/video";
import GetVideoRequest = video.GetVideoRequest;
import videoClient from "@/lib/grpcClient";
import UploadVideoRequest = video.UploadVideoRequest;
import ListVideosRequest = video.ListVideosRequest;
import { VideoListResponse } from "@/app/videos/_types/VideoListResponse";

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
 * Get video list
 * @param lastEvaluatedId
 */
export const getVideoList = async (lastEvaluatedId?: string) => {
  return new Promise<VideoListResponse>((resolve, reject) => {
    const request = new ListVideosRequest();
    if (lastEvaluatedId) {
      request.last_evaluated_key = lastEvaluatedId;
    }
    request.limit = 10;

    videoClient.ListVideos(request, {}, (err, response) => {
      if (err) {
        reject(err);
        return;
      }
      if (!response) {
        reject(new Error("No response"));
        return;
      }
      const videos = response.videos.map((video) => ({
        id: video.video_id,
        title: video.title,
        createdAt: video.created_at,
      }));
      const result: VideoListResponse = {
        videos,
        lastEvaluatedId: response._last_evaluated_key,
      };
      resolve(result);
    });
  });
};

/**
 * Upload video
 * @param fileBuffer
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
