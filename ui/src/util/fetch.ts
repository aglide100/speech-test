import axios from "axios";

export function getJobList(callback: any, limit?: number, offset?: number) {
    if (offset != undefined && offset < 0) {
        offset = 0;
    }

    // const destination = "https://jsonplaceholder.typicode.com/posts";
    const destination = "/data/list";
    axios
        .get(destination, { params: { limit: limit, offset: offset } })
        .then((res) => {
            callback(res);
        });
}

export function getJobText(callback: any, jobId: string) {
    const destination = "/data/text?jobId=" + jobId;
    axios.get(destination).then((res) => {
        callback(res);
    });
}
