import React from "react";
import { motion } from "framer-motion";
import { TailSpin } from "react-loader-spinner";

const transition = { duration: 1.5, ease: "easeInOut" };
export const Loading = () => {
    return (
        <motion.div
            // animate={{ x: 0 }}
            // initial={{ x: "100%" }}
            // exit={{ x: "-100%" }}
            transition={transition}
            className="flex flex-col w-full fixed z-30 inset-0 bg-gray-700 items-center justify-center"
        >
            <TailSpin
                height="80"
                width="80"
                color="#4fa94d"
                ariaLabel="tail-spin-loading"
                radius="1"
                wrapperStyle={{}}
                wrapperClass=""
                visible={true}
            />
            <div className="mt-10 text-white text-2xl flex flex-row">
                Loading component
                <div className="writer">
                    <div className="writer-text">...</div>
                </div>
            </div>
        </motion.div>
    );
};
