import React, { useRef } from "react";
import { motion } from "framer-motion";
import ReactHlsPlayer from "@ducanh2912/react-hls-player";

export interface ContentData {
    playListUrl: string;
    text: string;
    id: string;
}

export const Content = (props: ContentData) => {
    const playerRef = useRef<HTMLVideoElement | null>(null);

    return (
        <motion.div
            className="relative w-auto p-5 mt-10 text-center h-full"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.5, duration: 0.3 }}
        >
            <ReactHlsPlayer
                autoPlay={false}
                controls={true}
                width="100%"
                src={props.playListUrl}
                playerRef={playerRef}
                className="-mt-10 mb-5"
            ></ReactHlsPlayer>
            <motion.div
                layoutId={`card-${props.id}`}
                className="w-auto h-auto mt-5 text-left text-white text-base leading-loose font-mono"
            >
                {props.text}
            </motion.div>
        </motion.div>
    );
};
