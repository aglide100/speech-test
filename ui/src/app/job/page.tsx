"use client";
import React, { useEffect } from "react";

import { List } from "@/component/ItemList/ItemList";
import Item, { DataType } from "../../component/Item/Item";

import { useState } from "react";
import { getJobList } from "@/util/fetch";

import { AnimatePresence, motion } from "framer-motion";
import { Overlay } from "@/component/Overlay/Overlay";
export default function Home({
    params,
    searchParams,
}: {
    params: { slug: string };
    searchParams: { [key: string]: string | string[] | undefined };
}) {
    const [data, setData] = useState<DataType[]>([]);
    const [isLoading, setIsLoading] = useState(true);

    const [current, setCurrent] = useState<DataType>();

    useEffect(() => {
        if (isLoading) {
            getJobList((result: any) => {
                let list: DataType[] = [];

                result.data.map((d: any) => {
                    const tmp: DataType = {
                        Id: d.Id,
                        Content: d.Content,
                        PlayingTime: d.PlayingTime,
                        Speaker: d.Speaker,
                        background:
                            "#" + Math.random().toString(16).slice(2, 8),
                    };

                    list.push(tmp);
                });

                setData(list);
                setIsLoading(false);
            });
        }
    }, [isLoading]);

    return (
        <>
            {!isLoading && (
                <>
                    <AnimatePresence>
                        {current && (
                            <>
                                <Overlay
                                    handler={() => {
                                        setCurrent(undefined);
                                    }}
                                ></Overlay>
                                <Item
                                    data={current}
                                    handler={(id: string) => {
                                        setCurrent(undefined);
                                    }}
                                ></Item>
                            </>
                        )}
                    </AnimatePresence>
                    <List
                        items={data}
                        handler={(data: DataType) => {
                            setCurrent(data);
                        }}
                    />
                </>
            )}
        </>
    );
}
