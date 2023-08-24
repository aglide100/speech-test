"use client";
import React, { useState } from "react";
import axios from "axios";

const Test = () => {
    const [data, setData] = useState("");
    const [iframeUrl, setIframeUrl] = useState("");
    const [loadedUrl, setLoadedUrl] = useState("");

    const handleButtonClick = () => {
        // "Load" 버튼을 클릭하여 iframe의 src를 설정
        setLoadedUrl(iframeUrl);
        axios
            .get(iframeUrl) // 실제 데이터 가져오는 URL로 변경해야 함
            .then((response) => {
                setData(response.data);
            })
            .catch((error) => {
                console.log(error);
            });
    };

    return (
        <div>
            <h1>iframe</h1>
            <div>
                <input
                    type="text"
                    value={iframeUrl}
                    onChange={(e) => setIframeUrl(e.target.value)}
                    placeholder="Enter URL"
                />
                <button onClick={handleButtonClick}>Load</button>
            </div>
            <iframe
                title="Data Viewer"
                // srcDoc={data}
                src={loadedUrl} // 로드된 URL을 iframe의 src로 사용
                width="100%"
                height="400px"
                frameBorder="0"
            />
        </div>
    );
};

export default Test;
