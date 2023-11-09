"use client";
import { motion, useMotionValue } from "framer-motion";
import { Content as ContentBlock } from "../Content/Content";
import { closeSpring, openSpring } from "../../util/animation";
import { useRef, useEffect, useState } from "react";
import { useScrollConstraints } from "../../hook/useScrollConstraints";
import { useScrollEvent } from "../../hook/useScrollEvent";
import { getJobText } from "../../util/fetch";
import { Cover } from "../Cover/Cover";
import { Title } from "../Title/Title";

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
    const [dismissDistance, setDismissDistance] = useState(150);

    useEffect(() => {
        const handleResize = () => {
            if (window.innerWidth < 768) {
                setDismissDistance(100);
            } else {
                setDismissDistance(150);
            }
        };

        handleResize();

        window.addEventListener("resize", handleResize);

        return () => {
            window.removeEventListener("resize", handleResize);
        };
    }, []);

    const rootRef = useRef(null);
    const itemRef = useRef(null);

    const constraints = useScrollConstraints(itemRef, true);

    function checkSwipeToDismiss() {
        y.get() > dismissDistance && handler(data.Id);
    }

    useScrollEvent(rootRef, y, constraints, checkSwipeToDismiss, true);
    return (
        <div
            ref={rootRef}
            className="fixed top-0 left-0 right-0 w-screen md:w-auto h-auto z-20 overflow-hidden pt-10 md:p-10 pb-40"
        >
            <motion.div
                ref={itemRef}
                key={`card-container-${data.Id}`}
                layoutId={`card-container-${data.Id}`}
                initial={openSpring}
                exit={closeSpring}
                transition={{ duration: 0.5 }}
                drag={"y"}
                dragConstraints={constraints}
                className="z-50 relative rounded-lg overflow-hidden w-screen md:w-3/4 h-full mx-auto border-solid border-1 border-black shadow-xl"
                style={{ backgroundColor: "#1c1c1e", y: y }}
            >
                <div
                    className="text-white text-2xl absolute top-2 left-5 z-50"
                    onClick={(e) => {
                        e.preventDefault();
                        handler(data.Id);
                    }}
                >
                    {"<"}
                </div>
                <Cover isOped={true} background={data.background} />

                <motion.div
                    className="absolute top-4 left-11 font-bold"
                    layoutId={`title-container-${data.Id}`}
                >
                    <Title id={data.Id} title={data.Title} />
                </motion.div>

                <ContentBlock id={data.Id} playListUrl={m3u8Url} text={text} />
            </motion.div>
        </div>
    );
}
