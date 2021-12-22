const LoadFile = (props) => {
    console.log(props.filesURL);
    return (
        <>
            <h1>Uploaded Files</h1>
            <ul>
                {props.filesURL.map((fileURL) => (
                    <li>
                        <img src={fileURL} alt="hola" />
                    </li>
                ))}
            </ul>
        </>
    );
};

export default LoadFile;
