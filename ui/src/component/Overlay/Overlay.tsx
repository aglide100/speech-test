import React from "react";
import { motion } from "framer-motion";

export const Overlay = () => {
    return (
        <motion.div
            layoutId="overlay"
            exit={{
                WebkitBackdropFilter: "blur(0px)",
                backdropFilter: "blur(0px)",
            }}
            animate={{
                WebkitBackdropFilter: "blur(12px)",
                backdropFilter: "blur(12px)",
            }}
            transition={{ duration: 0.3, delay: 0.4 }}
            className="fixed top-0 left-0 bottom-0 w-full h-full z-20"
        ></motion.div>
    );
};
