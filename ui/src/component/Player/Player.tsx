import React, { useEffect, useRef } from "react";
import Hls from "hls.js";

interface HLSVideoPlayerProps {
    videoUrl: string;
}

const Player: React.FC<HLSVideoPlayerProps> = ({ videoUrl }) => {
    const videoRef = useRef<HTMLVideoElement | null>(null);

    useEffect(() => {
        if (videoRef.current) {
            const video = videoRef.current;

            if (Hls.isSupported()) {
                const hls = new Hls();
                hls.loadSource(videoUrl);
                hls.attachMedia(video);
                hls.on(Hls.Events.MANIFEST_PARSED, function () {
                    video.play();
                });
            } else if (video.canPlayType("application/vnd.apple.mpegurl")) {
                video.src = videoUrl;
                video.addEventListener("loadedmetadata", function () {
                    video.play();
                });
            }
        }
    }, [videoUrl]);

    return (
        <div>
            <video ref={videoRef} controls />
        </div>
    );
};

export default Player;
