import React, { useEffect, useRef } from "react";
import { motion } from "framer-motion";
import Hls from "hls.js";

export interface ContentData {
    playListUrl: string;
    text: string;
    id: string;
}

export const Content = (props: ContentData) => {
    const playerRef = useRef<HTMLVideoElement | null>(null);
    if (props.playListUrl.length == 0) {
        alert("can't get data, please check url");
    }

    useEffect(() => {
        const hls = new Hls();
        hls.loadSource(props.playListUrl + ".m3u8");
        if (playerRef.current != null) {
            hls.attachMedia(playerRef.current);
        }
    });
    return (
        <motion.div
            className="relative w-auto p-5 mt-10 text-center h-full"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.5, duration: 0.3 }}
        >
            <audio className="w-full" ref={playerRef} controls></audio>
            <motion.div
                layoutId={`card-${props.id}`}
                className="w-auto h-auto mt-5 text-left text-white text-base leading-loose font-mono"
            >
                {props.text}
            </motion.div>
        </motion.div>
    );
};
