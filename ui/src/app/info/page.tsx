"use client";
import React from "react";
import { motion, AnimatePresence } from "framer-motion";
import { initMain, exitMain, enterMain, transition } from "@/util/animation";

export default function Home({
    params,
    searchParams,
}: {
    params: { slug: string };
    searchParams: { [key: string]: string | string[] | undefined };
}) {
    return (
        <AnimatePresence>
            <motion.div
                key={"info-page"}
                initial={initMain}
                exit={exitMain}
                animate={enterMain}
                transition={transition}
            >
                <div>Info is here</div>
            </motion.div>
        </AnimatePresence>
    );
}
