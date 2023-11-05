import React, { useRef } from "react";
import { useRouter } from "next/navigation";
import { motion, useMotionValue } from "framer-motion";
import classNames from "classnames";
import Link from "next/link";
import ReactHlsPlayer from "@ducanh2912/react-hls-player";

export interface ContentData {
    id: string;
}

export const Content = (props: ContentData) => {
    const playerRef = useRef<HTMLVideoElement | null>(null);

    return (
        <motion.div className="m-auto text-center ">
            <ReactHlsPlayer
                autoPlay={false}
                controls={true}
                width="100%"
                height="auto"
                src={"/data/playlist/39"}
                // src={"/data/playlist/" + props.id}
                playerRef={playerRef}
            ></ReactHlsPlayer>
            <motion.div
                layoutId={`card-${props.id}`}
                className="w-auto h-auto mt-5 text-white"
            >
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                Pellentesque quis volutpat tortor. In hac habitasse platea
                dictumst. Vivamus eget ante dignissim, suscipit sem ut, pharetra
                sapien. Ut at erat magna. Nam ac maximus odio, nec eleifend
                turpis. Donec facilisis consequat lorem, quis sodales justo
                commodo mollis. Vivamus placerat id elit ac convallis. Nam at
                sem vulputate, tincidunt ligula sed, dapibus nisi. In a nulla
                diam. Integer fermentum blandit nisl. Duis est mauris,
                sollicitudin non risus at, tempor porta odio. Nunc ut justo
                scelerisque, hendrerit dui et, Lorem ipsum dolor sit amet,
                consectetur adipiscing elit. Pellentesque quis volutpat tortor.
                In hac habitasse platea dictumst. Vivamus eget ante dignissim,
                suscipit sem ut, pharetra sapien. Ut at erat magna. Nam ac
                maximus odio, nec eleifend turpis. Donec facilisis consequat
                lorem, quis sodales justo commodo mollis. Vivamus placerat id
                elit ac convallis. Nam at sem vulputate, tincidunt ligula sed,
                dapibus nisi. In a nulla diam. Integer fermentum blandit nisl.
                Duis est mauris, sollicitudin non risus at, tempor porta odio.
                Nunc ut justo scelerisque, hendrerit dui et,
            </motion.div>
        </motion.div>
    );
};
