"use client";
import { AnimatePresence, LayoutGroup } from "framer-motion";

export default function JobLayout({ children }: { children: React.ReactNode }) {
    return (
        <div className="p-3 md:p-20">
            <LayoutGroup>
                <AnimatePresence>{children}</AnimatePresence>
            </LayoutGroup>
        </div>
    );
}
