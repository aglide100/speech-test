import React, { useRef } from "react";
import { motion } from "framer-motion";
import ReactHlsPlayer from "@ducanh2912/react-hls-player";

export interface ContentData {
    id: string;
    data: string;
}

export const Content = (props: ContentData) => {
    const playerRef = useRef<HTMLVideoElement | null>(null);

    return (
        <motion.div className="m-auto text-center">
            <ReactHlsPlayer
                autoPlay={false}
                controls={true}
                width="100%"
                height="auto"
                src={`/data/playlist/${props.id}`}
                playerRef={playerRef}
            ></ReactHlsPlayer>
            <motion.div
                layoutId={`card-${props.id}`}
                className="w-auto h-auto mt-5 text-left text-white leading-relaxed"
            >
                {props.data}
            </motion.div>
        </motion.div>
    );
};
