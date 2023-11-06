import React from "react";
import { motion } from "framer-motion";
import classNames from "classnames";

export interface CardProps {
    id: string;
    idx: number;
    title: string;
    category: string;
    background: string;
    handler(
        id: string,
        title: string,
        category: string,
        background: string
    ): void;
}

function Card({ id, idx, title, category, background, handler }: CardProps) {
    return (
        <motion.li
            layoutId={`Item-${id}`}
            className={classNames(
                "relative p-6 h-72 w-full basis-full md:w-2/4",
                idx % 2 == 0 ? "md:pr-0" : "md:pl-0",
                idx % 4 == 1 || idx % 4 == 0
                    ? "flex-grow-0 flex-shrink-0 md:basis-3/5"
                    : "flex-grow-0 flex-shrink-0 md:basis-2/5"
            )}
            onClick={(e) => {
                e.preventDefault();
                handler(id, title, category, background);
            }}
        >
            <div className="w-full h-full relative block pointer-events-none">
                <motion.div
                    className="pointer-events-auto relative rounded-lg bg-gray-900 overflow-hidden w-full h-full mx-auto"
                    style={{ backgroundColor: "#1c1c1e" }}
                    layoutId={`card-container-${id}`}
                >
                    <motion.div
                        className="absolute top-0 left-0 overflow-hidden w-full h-full"
                        layoutId={`card-image-container-${id}`}
                    >
                        <div
                            className="object-cover h-full w-full"
                            // style={{ backgroundColor: background }}
                        ></div>
                        {/* <img
                            className="object-cover h-full w-full"
                            src={`/images/${id}.jpg`}
                            alt=""
                        /> */}
                    </motion.div>
                    <motion.div
                        className="absolute left-4 top-4"
                        layoutId={`title-container-${id}`}
                    >
                        <span className="text-white text-2xl uppercase">
                            {category}
                        </span>
                        <h2 className="text-white text-xl my-2">{title}</h2>
                    </motion.div>
                </motion.div>
            </div>
        </motion.li>
    );
}

export interface ListProps {
    items: any[];
    handler(id: string, title: string, category: string, bg: string): void;
}

export function List(props: ListProps) {
    return (
        <ul className="list-none m-0 p-0 flex flex-wrap content-start">
            {props.items.map((card, idx) => (
                <Card
                    key={"key_card_" + card.id}
                    {...card}
                    idx={idx + 1}
                    background={"#" + Math.random().toString(16).slice(2, 8)}
                    handler={props.handler}
                />
            ))}
        </ul>
    );
}
