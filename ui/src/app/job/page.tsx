"use client";
import React from "react";

import { List } from "@/component/ItemList/ItemList";
import Item from "../../component/Item/Item";

import { useState } from "react";
const items: any[] = [
    {
        id: "1",
        category: "",
        title: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
    },
    {
        id: "2",
        category: "",
        title: "Quisque sed interdum ligula, ut ullamcorper libero",
    },
    {
        id: "3",
        category: "",
        title: "Praesent finibus est non est pellentesque placerat",
    },
    {
        id: "4",
        category: "",
        title: "Pellentesque pharetra scelerisque sem non blandit",
    },
    {
        id: "5",
        category: "",
        title: "Mauris eu mauris fringilla",
    },
    {
        id: "6",
        category: "",
        title: "ccumsan lorem eu",
    },
];

export default function Home({
    params,
    searchParams,
}: {
    params: { slug: string };
    searchParams: { [key: string]: string | string[] | undefined };
}) {
    const [current, setCurrent] = useState("");
    const [category, setCategory] = useState("");
    const [title, setTitle] = useState("");
    const [background, setBackground] = useState("");

    return (
        <>
            {current != "" ? (
                <>
                    <Item
                        id={current}
                        category={category}
                        title={title}
                        background={background}
                        handler={(id: string) => {
                            setCurrent("");
                        }}
                    ></Item>
                </>
            ) : (
                <></>
            )}
            <List
                items={items}
                handler={(id: string, t: string, c: string, bg: string) => {
                    setCurrent(id);
                    setTitle(t);
                    setCategory(c);
                    setBackground(bg);
                }}
            />
        </>
    );
}
