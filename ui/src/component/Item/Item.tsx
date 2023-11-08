"use client";
import { AnimatePresence, motion, useMotionValue } from "framer-motion";
import { Content as ContentBlock } from "../Content/Content";
import { closeSpring, openSpring } from "@/hook/animation";
import { useRef, useEffect, useState } from "react";
import { useScrollConstraints } from "@/hook/useScrollConstraints";
import { useWheelScroll } from "@/hook/useWheelScroll";
import { getJobText } from "@/util/fetch";

export interface DataType {
    Id: string;
    Speaker: string;
    PlayingTime: Float32Array;
    Title: string;
    background: string;
}

export interface FrameProps {
    data: DataType;
    handler(id: string): void;
}

interface ItemProps extends FrameProps {
    m3u8Url: string;
    text: string;
}

const dismissDistance = 150;

export default function Frame({ data, handler }: FrameProps) {
    const [isLoading, setIsLoading] = useState(true);
    const [playListUrl, setPlayListUrl] = useState("");
    const [text, setText] = useState("");

    useEffect(() => {
        if (isLoading) {
            setPlayListUrl(`/data/playlist?jobId=${data.Id}`);

            getJobText((result: any) => {
                setText(result.data.Txt);
                setIsLoading(false);
            }, data.Id);
        }
    }, [isLoading]);

    return (
        <motion.div className="fixed z-20" layoutId={`card-root-${data.Id}`}>
            {!isLoading ? (
                <Item
                    data={data}
                    handler={handler}
                    m3u8Url={playListUrl}
                    text={text}
                />
            ) : (
                <></>
            )}
        </motion.div>
    );
}

function Item({ data, handler, m3u8Url, text }: ItemProps) {
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
            className="fixed top-0 left-0 right-0 w-screen md:w-auto h-auto z-20 overflow-hidden pt-10 md:p-10 pb-40"
            // initial={{ opacity: 0 }}
            // exit={{ opacity: 0 }}
            // animate={{ opacity: 1 }}
        >
            <motion.div
                ref={itemRef}
                layoutId={`card-container-${data.Id}`}
                initial={openSpring}
                exit={closeSpring}
                drag={"y"}
                dragConstraints={constraints}
                className="z-50 relative rounded-lg overflow-hidden w-screen md:w-3/4 h-full mx-auto border-solid border-1 border-black shadow-xl"
                style={{ backgroundColor: "#1c1c1e", y: y }}
            >
                <motion.div
                    className="md:relative top-0 left-0 overflow-hidden w-screen md:w-full"
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
                            {data.Title}
                        </p>
                    </motion.div>
                </motion.div>

                <motion.div className="relative z-30 w-auto h-auto p-5  mt-10">
                    <ContentBlock
                        id={data.Id}
                        playListUrl={m3u8Url}
                        text={text}
                    />
                </motion.div>
            </motion.div>
        </motion.div>
    );
}
