"use client";
import React from "react";
import { AnimatePresence, motion } from "framer-motion";
import { List } from "@/component/ItemList/ItemList";
export default function Home({
    params,
    searchParams,
}: {
    params: { slug: string };
    searchParams: { [key: string]: string | string[] | undefined };
}) {
    // const router = useRouter();
    // router.push("/list");

    return (
        <AnimatePresence>
            <List></List>
            {/* <motion.div
                key={"main-page"}
                initial={initMain}
                exit={exitMain}
                animate={enterMain}
                transition={transition}
                className="w-full h-full flex justify-center text-center"
            >
                <button className="mr-10">
                    <Link href={"/info"}>Info</Link>
                </button>
                <button>
                    <Link href={"/list"}>list</Link>
                </button>
            </motion.div> */}
        </AnimatePresence>
    );
}
