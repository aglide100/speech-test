"use client";
import { Header } from "@/component/Header/Header";
import "./globals.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { AnimatePresence, LayoutGroup } from "framer-motion";

const inter = Inter({ subsets: ["latin"] });

const metadata: Metadata = {
    title: " ",
    description: " ",
};

export default function RootLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <html>
            <body>
                <Header />
                <div className="md:p-20">{children}</div>
            </body>
        </html>
    );
}
