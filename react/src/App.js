import ReactDOM from "react-dom";
import UploadFile from "./components/uploadfile";
import Load from "./components/loadfile";

const App = () => {
    let handleUpload = () => {
        ReactDOM.render(<UploadFile />, document.getElementById("root"));
    };
    let handleLoad = () => {
        ReactDOM.render(<Load />, document.getElementById("root"));
    };
    return (
        <>
            <h1>Welcome</h1>
            <button onClick={handleUpload}>Upload File</button>
            <button onClick={handleLoad}>Load File</button>
        </>
    );
};

export default App;
