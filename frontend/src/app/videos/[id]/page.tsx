"use client";
import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import {getVideoURL} from "@/app/videos/action";

export default function Page() {
    const { id } = useParams();
    const [videoURL, setVideoURL] = useState<string | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        if (id) {
            const _ = fetchVideo();
        }
    }, [id]);

    const fetchVideo = async () => {
        try {
            const url = await getVideoURL(id as string);
            setVideoURL(url);
        } catch (error) {
            console.error("Failed to load video", error);
        } finally {
            setLoading(false);
        }
    };

    if (loading) return <div className="p-4">Loading...</div>;

    return (
        <div className="p-4">
            <h1 className="text-2xl font-bold">Video Player</h1>
            {videoURL ? (
                <video controls className="mt-4 w-full max-w-lg">
                    <source src={videoURL} type="video/mp4" />
                    Your browser does not support the video tag.
                </video>
            ) : (
                <p className="text-red-500">Failed to load video</p>
            )}
        </div>
    );
}
