import Image from "next/image";
import Player from "@/component/\bPlayer/Player";

export default function Home() {
    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex"></div>
            <Player videoUrl="localhost:8080/test"></Player>
        </main>
    );
}
