import React, { ReactNode } from "react";
import { motion } from "framer-motion";

const exitOverlay = {
    WebkitBackdropFilter: "blur(0px)",
    backdropFilter: "blur(0px)",
    zIndex: 0,
};

const openOverlay = {
    WebkitBackdropFilter: "blur(12px)",
    backdropFilter: "blur(12px)",
};

export type OverlayProps = {
    isOpen: boolean;
    children?: ReactNode;
};

export const Overlay = (props: OverlayProps) => {
    return (
        <motion.div
            animate={props.isOpen ? openOverlay : exitOverlay}
            transition={{ duration: 0.4, delay: props.isOpen ? 0.4 : 0 }}
            className="fixed top-0 left-0 bottom-0 w-full h-full z-10"
        >
            {props.children}
        </motion.div>
    );
};
