"use client";
import { motion, useMotionValue } from "framer-motion";
import { Content } from "../Content/Content";
import { closeSpring, openSpring } from "@/hook/animation";
import { useRef } from "react";
import { useScrollConstraints } from "@/hook/useScrollConstraints";
import { useWheelScroll } from "@/hook/useWheelScroll";

export interface ItemProps {
    id: string;
    category: string;
    title: string;
    background: string;
    handler(id: string): void;
}

const dismissDistance = 150;

export default function Item({
    id,
    category,
    title,
    background,
    handler,
}: ItemProps) {
    const y = useMotionValue(0);

    const rootRef = useRef(null);
    const itemRef = useRef(null);

    const constraints = useScrollConstraints(itemRef, true);

    function checkSwipeToDismiss() {
        y.get() > dismissDistance && handler(id);
    }

    useWheelScroll(rootRef, y, constraints, checkSwipeToDismiss, true);
    return (
        <motion.div ref={rootRef}>
            <motion.div
                onClick={(e) => {
                    e.preventDefault();
                    handler(id);
                }}
                layoutId="overlay"
                exit={{ opacity: 1 }}
                initial={{ opacity: 0 }}
                animate={{ opacity: 0.75, backgroundColor: "black" }}
                transition={{ duration: 0.3 }}
                className="fixed top-0 bottom-0 w-full h-full z-10"
            ></motion.div>
            <motion.div className="fixed top-0 left-0 right-0 z-20 overflow-hidden p-10">
                <motion.div
                    ref={itemRef}
                    layoutId={`card-container-${id}`}
                    initial={openSpring}
                    exit={closeSpring}
                    drag={"y"}
                    dragConstraints={{ top: -10, bottom: 10 }}
                    className="pointer-events-auto relative rounded-lg overflow-hidden w-full md:w-3/4 h-full mx-auto"
                    style={{ backgroundColor: "#1c1c1e", y: y }}
                >
                    <motion.div
                        className="relative top-0 left-0 overflow-hidden w-full"
                        layoutId={`card-image-container-${id}`}
                    >
                        <div
                            className="text-white text-2xl absolute top-2 left-5 z-20"
                            onClick={(e) => {
                                e.preventDefault();
                                handler(id);
                            }}
                        >
                            {"<"}
                        </div>
                        <div
                            className="w-full h-80"
                            // style={{ backgroundColor: background }}
                        ></div>
                        {/* <img
                            className="w-full h-80"
                            src={`/images/${id}.jpg`}
                            alt=""
                        /> */}

                        <motion.div
                            className=" absolute top-3 left-11"
                            layoutId={`title-container-${id}`}
                        >
                            <span className="text-white text-2xl uppercase">
                                {category}
                            </span>
                            <h2 className="text-white text-xl my-2">{title}</h2>
                        </motion.div>
                    </motion.div>

                    <motion.div className="relative z-30 w-auto h-auto p-5  mt-10">
                        <Content id={id} />
                    </motion.div>
                </motion.div>
            </motion.div>
        </motion.div>
    );
}
