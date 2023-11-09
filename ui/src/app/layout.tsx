"use client";
import { AnimatePresence, LayoutGroup } from "framer-motion";
import { Header } from "../component/Header/Header";
import "./globals.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";

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
                <main>{children}</main>
            </body>
        </html>
    );
}
