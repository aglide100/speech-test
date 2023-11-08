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
        <motion.div className="m-auto text-center h-full">
            <ReactHlsPlayer
                autoPlay={false}
                controls={true}
                width="100%"
                src={props.playListUrl}
                playerRef={playerRef}
                className="h-10 -mt-10 mb-5"
            ></ReactHlsPlayer>
            <motion.div
                layoutId={`card-${props.id}`}
                className="w-auto h-auto mt-5 text-left text-white text-base leading-loose font-mono"
            >
                {/* {text} */}
                {props.text}
            </motion.div>
        </motion.div>
    );
};
