import ReactDOM from "react-dom";
import UploadFile from "./components/uploadfile";
import LoadFile from "./components/loadfile";

const App = () => {
    let handleUpload = () => {
        ReactDOM.render(<UploadFile />, document.getElementById("root"));
    };
    let handleLoad = () => {
        fetch("http://localhost:8080/api/v1/load", {})
            .then((responsive) => {
                if (responsive.status >= 400) {
                    throw responsive.status;
                }
                return responsive.json();
            })
            .then((resp) => {
                console.log(resp);
                ReactDOM.render(<LoadFile />, document.getElementById("root"));
            })
            .catch(() => {
                ReactDOM.render(
                    <h1>Failure to Load File</h1>,
                    document.getElementById("root")
                );
            });
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
