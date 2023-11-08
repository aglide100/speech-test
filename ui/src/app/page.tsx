"use client";
import React from "react";

import { List } from "@/component/ItemList/ItemList";

export default function Home({
    params,
    searchParams,
}: {
    params: { slug: string };
    searchParams: { [key: string]: string | string[] | undefined };
}) {
    return (
        <>
            <List></List>
        </>
    );
}
