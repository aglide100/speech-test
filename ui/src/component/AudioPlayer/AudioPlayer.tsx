import React, { useRef, useState } from "react";

function AudioPlayer(url: any) {
    const [isPlaying, setIsPlaying] = useState(false);

    const audioRef = useRef();

    const togglePlay = () => {
        const audio = new Audio(url);
        if (isPlaying) {
            audio.pause();
        } else {
            audio.play();
        }
        setIsPlaying(!isPlaying);
    };

    return (
        <div>
            <audio src={url}></audio>
            <button onClick={togglePlay}>
                {isPlaying ? "일시정지" : "재생"}
            </button>
        </div>
    );
}

export default AudioPlayer;
