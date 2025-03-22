export type VideoResponse = {
    id: string;
    title: string;
    createdAt: string;
}

export type VideoListResponse = {
    videos: VideoResponse[];
    lastEvaluatedId?: string;
}
