import React, { useEffect, useState } from "react";

import classNames from "classnames";

import { getJobList } from "../../util/fetch";
import { AnimatePresence, motion, LayoutGroup } from "framer-motion";
import { Overlay } from "../../component/Overlay/Overlay";
import { useQueryState } from "next-usequerystate";
import { Loading } from "../../component/Loading/Loading";
import InfiniteScroll from "react-infinite-scroll-component";
import Item, { DataType } from "../../component/Item/Item";
import { Cover } from "../Cover/Cover";
import { Title } from "../Title/Title";

export interface CardProps {
    idx: number;
    data: DataType;
    handler(data: DataType): void;
}

function Card({ data, idx, handler }: CardProps) {
    return (
        <motion.li
            layoutId={`card-root-${data.Id}`}
            className={classNames(
                "relative p-6 h-72 w-full basis-full md:w-1/4 z-0",
                idx % 2 == 0 ? "md:pr-0" : "md:pl-0",
                idx % 4 == 1 || idx % 4 == 0
                    ? "flex-grow-0 flex-shrink-0 md:basis-3/5"
                    : "flex-grow-0 flex-shrink-0 md:basis-2/5"
            )}
            // style={{ zIndex: 10 }}
            onClick={(e) => {
                e.preventDefault();
                handler(data);
            }}
        >
            <div className="w-full h-full relative block pointer-events-none">
                <motion.div
                    className="pointer-events-auto relative rounded-lg shadow-2xl  overflow-hidden w-full h-full mx-auto"
                    style={{ backgroundColor: "#1c1c1e" }}
                    layoutId={`card-container-${data.Id}`}
                >
                    <Cover isOped={false} background={data.background} />
                </motion.div>
                <motion.div
                    className="absolute left-4 top-4 font-bold"
                    layoutId={`title-container-${data.Id}`}
                >
                    <Title id={data.Id} title={data.Title} />
                </motion.div>
            </div>
        </motion.li>
    );
}

const limit = 6;

export function List() {
    const [data, setData] = useState<DataType[]>([]);
    const [page, setPage] = useState(0);
    const [memo, setMemo] = useState<number[]>([]);

    const [isLoading, setIsLoading] = useState(true);
    const [isLast, setIsLast] = useState(false);
    const [current, setCurrent] = useState<DataType>();

    const [job, setJob] = useQueryState("job", { shallow: true });

    const fetchMore = (callback?: any) => {
        let ok = true;

        memo.forEach((m) => {
            if (m == page) {
                ok = false;
            }
        });

        if (ok) {
            getJobList(
                (result: any) => {
                    let list: DataType[] = [];

                    if (result.data == undefined) {
                        callback();
                    } else if (result.data.length != 0) {
                        if (result.data.length < limit) {
                            setIsLast(true);
                        }

                        result.data.map((d: any) => {
                            const tmp: DataType = {
                                Id: d.Id,
                                Title: d.Title,
                                PlayingTime: d.PlayingTime,
                                Speaker: d.Speaker,
                                background:
                                    "#" +
                                    Math.random().toString(16).slice(2, 8),
                            };

                            list.push(tmp);
                        });

                        setData((d) => d.concat(list));
                        setMemo((m) => m.concat(page));
                        setPage((p) => p + 1);
                        callback();
                    }
                },
                limit,
                page * limit - limit
            );
        }
    };

    useEffect(() => {
        if (isLoading) {
            fetchMore(() => {
                setIsLoading(false);
            });
        }

        if (job != undefined && !isLoading) {
            data.forEach((val) => {
                if (val.Id == job) {
                    setCurrent(val);
                    return;
                }
            });
        }
    }, [isLoading]);

    const openHandler = (data: DataType) => {
        setCurrent(data);
    };
    const closeHandler = () => {
        setCurrent(undefined);
        if (job != undefined) {
            setJob(null);
        }
    };
    return (
        <AnimatePresence
            mode="wait"
            // onExitComplete={() => {
            //     if (current != undefined) {
            //         setJob(current.Id);
            //     } else {
            //         setJob(null);
            //     }
            // }}
        >
            <LayoutGroup>
                {!isLoading ? (
                    <motion.div
                        animate={{ opacity: 1 }}
                        initial={{ opacity: 0 }}
                    >
                        <InfiniteScroll
                            dataLength={data.length}
                            next={() => fetchMore()}
                            hasMore={!isLast}
                            loader={<></>}
                            style={{
                                zIndex: 10,
                                height: "100%",
                                width: "100%",
                                position: "absolute",
                                left: 0,
                                top: 0,
                            }}
                        >
                            <Overlay isOpen={current == null ? false : true} />
                            {current && (
                                <Item
                                    data={current}
                                    handler={(id: string) => {
                                        closeHandler();
                                    }}
                                ></Item>
                            )}

                            <motion.ul className="relative top-20 list-none md:p-20 p-0 flex flex-wrap content-start md:-mt-10 mt-10">
                                {data.map((card, idx) => (
                                    <Card
                                        key={"key__card_" + card.Id}
                                        data={card}
                                        idx={idx + 1}
                                        handler={openHandler}
                                    />
                                ))}
                            </motion.ul>
                        </InfiniteScroll>
                    </motion.div>
                ) : (
                    <motion.div
                        key="loading_component"
                        animate={{ opacity: 1 }}
                        initial={{ opacity: 0 }}
                        exit={{ opacity: 0 }}
                        transition={{ type: "spring", duration: 0.5 }}
                    >
                        <Loading key={"loading"} />
                    </motion.div>
                )}
            </LayoutGroup>
        </AnimatePresence>
    );
}
