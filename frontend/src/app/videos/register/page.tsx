"use client";
import { useState } from "react";
import { uploadVideo } from "@/app/videos/action";

/**
 * Video upload page
 * @constructor
 */
export default function Page() {
  const [title, setTitle] = useState("");
  const [file, setFile] = useState<File | null>(null);
  const [isUploading, setIsUploading] = useState(false);
  const [message, setMessage] = useState("");

  const handleUpload = async () => {
    if (!title || !file) {
      setMessage("タイトルとファイルを選択してください");
      return;
    }
    setIsUploading(true);
    setMessage("アップロード中...");

    try {
      const buffer = await file.arrayBuffer();
      const videoId = await uploadVideo(buffer, title);
      setMessage(`アップロード成功！動画ID: ${videoId}`);
    } catch (error) {
      setMessage("アップロード失敗: " + (error as Error).message);
    } finally {
      setIsUploading(false);
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-4 border rounded-lg shadow-lg">
      <h2 className="text-xl font-bold mb-4">動画アップロード</h2>
      <input
        type="text"
        placeholder="動画タイトル"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        className="w-full px-3 py-2 border rounded mb-3"
      />
      <input
        type="file"
        accept=".mp4"
        onChange={(e) => setFile(e.target.files?.[0] || null)}
        className="w-full px-3 py-2 border rounded mb-3"
      />
      <button
        onClick={handleUpload}
        disabled={isUploading}
        className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600 disabled:bg-gray-400"
      >
        {isUploading ? "アップロード中..." : "アップロード"}
      </button>
      {message && <p className="mt-3 text-sm text-center">{message}</p>}
    </div>
  );
}
