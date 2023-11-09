import { closeSpring, openSpring } from "@/util/animation";
import { motion } from "framer-motion";
import React from "react";

export interface CoverProps {
    background: string;
    isOped: boolean;
}

export const Cover = (props: CoverProps) => {
    return (
        <motion.div
            className="block top-0 left-0 overflow-hidden w-screen h-80"
            layoutId={"card-cover"}
            style={{ originX: 0, originY: 0 }}
        >
            <motion.img
                className="object-cover"
                alt=""
                src="https://picsum.photos/200"
                initial={false}
                animate={props.isOped ? { x: 0, y: -20 } : { x: 0, y: 0 }}
                transition={closeSpring}
                style={{
                    // backgroundColor: props.background,
                    width: "120%",
                    height: "120%",
                }}
            />
        </motion.div>
    );
};
