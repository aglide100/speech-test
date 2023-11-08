import React from "react";
import { motion } from "framer-motion";

export interface OverlayProps {
    // handler(): void;
}

export const Overlay = (props: OverlayProps) => {
    return (
        <motion.div
            // onClick={(e) => {
            //     e.preventDefault();
            //     props.handler();
            // }}
            layoutId="overlay"
            exit={{
                WebkitBackdropFilter: "blur(0px)",
                backdropFilter: "blur(0px)",
            }}
            animate={{
                WebkitBackdropFilter: "blur(12px)",
                backdropFilter: "blur(12px)",
            }}
            transition={{ duration: 0.3 }}
            className="fixed top-0 left-0 bottom-0 w-full h-full z-20"
        ></motion.div>
    );
};
