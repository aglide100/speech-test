export interface TitleData {
    id: string;
    title: string;
}

export const Title = (props: TitleData) => {
    return (
        <>
            <span className="text-white text-base uppercase">{props.id}</span>
            <p className="text-white text-2xl my-2 line-clamp-1">
                {props.title}
            </p>
        </>
    );
};
