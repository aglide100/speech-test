"use client";
import Image from "next/image";
import React, { useRef, useState } from "react";
import ReactHlsPlayer from "react-hls-player";

export default function Home() {
    const playerRef = useRef<HTMLVideoElement | null>(null);
    const [index, setIndex] = useState("39");
    const [url, setUrl] = useState("/data/playlist/" + index);

    const onChange = (e: any) => {
        setIndex(e.target.value);
    };

    const whenClick = () => {
        setUrl("/data/playlist/" + index);
    };

    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex"></div>
            <div>
                <input
                    onChange={onChange}
                    value={index}
                    className="text-black"
                />
                <button onClick={whenClick}>set</button>
            </div>
            <ReactHlsPlayer
                autoPlay={false}
                controls={true}
                width="100%"
                height="auto"
                src={url}
                playerRef={playerRef}
            ></ReactHlsPlayer>
        </main>
    );
}
