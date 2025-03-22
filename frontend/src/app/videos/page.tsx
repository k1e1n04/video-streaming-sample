"use client";
import { useEffect, useState, useRef } from "react";
import Link from "next/link";
import { getVideoList } from "@/app/videos/action";

export default function Page() {
    const [videos, setVideos] = useState<{ id: string; title: string }[]>([]);
    const [lastEvaluatedId, setLastEvaluatedId] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);
    const hasFetched = useRef(false); // Prevent fetching data on initial render

    useEffect(() => {
        if (!hasFetched.current) {
            hasFetched.current = true;
            fetchVideos();
        }
    }, []);

    const fetchVideos = async (lastId?: string) => {
        if (loading) return;
        setLoading(true);

        try {
            const response = await getVideoList(lastId);
            setVideos((prev) => {
                const newVideos = response.videos.filter((v) => !prev.some((p) => p.id === v.id));
                return [...prev, ...newVideos];
            });
            setLastEvaluatedId(response.lastEvaluatedId || null);
        } catch (error) {
            console.error("Failed to load videos", error);
        }

        setLoading(false);
    };

    return (
        <div className="p-4">
            <h1 className="text-2xl font-bold">Video List</h1>
            <ul className="mt-4 space-y-2">
                {videos.map((video) => (
                    <li key={video.id}>
                        <Link href={`/videos/${video.id}`} className="text-blue-500 hover:underline">
                            {video.title}
                        </Link>
                    </li>
                ))}
            </ul>
            {lastEvaluatedId && (
                <button
                    className="mt-4 bg-blue-500 text-white px-4 py-2 rounded"
                    onClick={() => fetchVideos(lastEvaluatedId)}
                    disabled={loading}
                >
                    {loading ? "Loading..." : "Load More"}
                </button>
            )}
        </div>
    );
}
