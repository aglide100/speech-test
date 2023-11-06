"use client";
import { motion, useMotionValue } from "framer-motion";
import { Content as ContentBlock } from "../Content/Content";
import { closeSpring, openSpring } from "@/hook/animation";
import { useRef } from "react";
import { useScrollConstraints } from "@/hook/useScrollConstraints";
import { useWheelScroll } from "@/hook/useWheelScroll";

export interface DataType {
    Id: string;
    Speaker: string;
    PlayingTime: Float32Array;
    Content: string;
    background: string;
}

export interface ItemProps {
    data: DataType;
    handler(id: string): void;
}

const dismissDistance = 150;

export default function Item({ data, handler }: ItemProps) {
    const y = useMotionValue(0);

    const rootRef = useRef(null);
    const itemRef = useRef(null);

    const constraints = useScrollConstraints(itemRef, true);

    function checkSwipeToDismiss() {
        y.get() > dismissDistance && handler(data.Id);
    }

    useWheelScroll(rootRef, y, constraints, checkSwipeToDismiss, true);
    return (
        <motion.div
            ref={rootRef}
            className="fixed top-0 left-0 right-0 z-20 w-full  overflow-hidden p-10"
        >
            <motion.div
                ref={itemRef}
                layoutId={`card-container-${data.Id}`}
                initial={openSpring}
                exit={closeSpring}
                drag={"y"}
                dragConstraints={{ top: -30, bottom: 30 }}
                className="pointer-events-auto relative rounded-lg overflow-hidden w-full md:w-3/4 h-full mx-auto border-solid border-1 border-black shadow-xl"
                style={{ backgroundColor: "#1c1c1e", y: y }}
            >
                <motion.div
                    className="relative top-0 left-0 overflow-hidden w-full"
                    layoutId={`card-image-container-${data.Id}`}
                >
                    <div
                        className="text-white text-2xl absolute top-2 left-5 z-20"
                        onClick={(e) => {
                            e.preventDefault();
                            handler(data.Id);
                        }}
                    >
                        {"<"}
                    </div>
                    <div
                        className="w-full h-80"
                        style={{ backgroundColor: data.background }}
                    ></div>
                    {/* <img
                            className="w-full h-80"
                            src={`/images/${id}.jpg`}
                            alt=""
                        /> */}

                    <motion.div
                        className="absolute top-4 left-11"
                        layoutId={`title-container-${data.Id}`}
                    >
                        <span className="text-white text-base uppercase">
                            {data.Id}
                        </span>
                        <p className="text-white text-2xl my-2 line-clamp-1">
                            {data.Content}
                        </p>
                    </motion.div>
                </motion.div>

                <motion.div className="relative z-30 w-auto h-auto p-5  mt-10">
                    <ContentBlock id={data.Id} data={data.Content} />
                </motion.div>
            </motion.div>
        </motion.div>
    );
}
