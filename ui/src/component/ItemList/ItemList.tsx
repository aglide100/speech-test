import React from "react";
import { motion } from "framer-motion";
import classNames from "classnames";
import { DataType } from "../Item/Item";

export interface CardProps {
    idx: number;
    data: DataType;
    handler(data: DataType): void;
}

function Card({ data, idx, handler }: CardProps) {
    return (
        <motion.li
            layoutId={`Item-${data.Id}`}
            className={classNames(
                "relative p-6 h-72 w-full basis-full md:w-2/4 ",
                idx % 2 == 0 ? "md:pr-0" : "md:pl-0",
                idx % 4 == 1 || idx % 4 == 0
                    ? "flex-grow-0 flex-shrink-0 md:basis-3/5"
                    : "flex-grow-0 flex-shrink-0 md:basis-2/5"
            )}
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
                    <motion.div
                        className="absolute top-0 left-0 overflow-hidden w-full h-full"
                        layoutId={`card-image-container-${data.Id}`}
                    >
                        <div
                            className="object-cover h-full w-full"
                            style={{ backgroundColor: data.background }}
                        ></div>
                        {/* <img
                            className="object-cover h-full w-full"
                            src={`/images/${id}.jpg`}
                            alt=""
                        /> */}
                    </motion.div>
                    <motion.div
                        className="absolute left-4 top-4"
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
            </div>
        </motion.li>
    );
}

export interface ListProps {
    items: DataType[];
    handler(data: DataType): void;
}

export function List(props: ListProps) {
    return (
        <ul className="list-none m-0 p-0 flex flex-wrap content-start">
            {props.items.map((card, idx) => (
                <Card
                    key={"key_card_" + card.Id}
                    data={card}
                    idx={idx + 1}
                    handler={props.handler}
                />
            ))}
        </ul>
    );
}
