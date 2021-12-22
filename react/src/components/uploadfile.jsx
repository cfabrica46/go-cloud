import ReactDOM from "react-dom";
import React, { useState } from "react";
const UploadFile = () => {
    const [selectedFile, setSelectedFile] = useState(undefined);

    const handleSubmit = () => {
        let dataForm = new FormData();
        dataForm.append("file", selectedFile);

        fetch("http://localhost:8080/api/v1/upload", {
            method: "POST",
            body: dataForm,
        })
            .then((responsive) => {
                if (responsive.status >= 400) {
                    console.log(responsive.status);
                    throw responsive.status;
                }
            })
            .then(() => {
                ReactDOM.render(
                    <h1>Uploaded File</h1>,
                    document.getElementById("root")
                );
            })
            .catch(() => {
                ReactDOM.render(
                    <h1>Failure to Upload File</h1>,
                    document.getElementById("root")
                );
            });
    };

    return (
        <>
            <h1>Upload File</h1>
            <input
                name="file"
                type="file"
                accept="image/png, image/jpeg"
                onChange={(e) => setSelectedFile(e.target.files[0])}
            />
            <button onClick={handleSubmit}>Upload</button>
        </>
    );
};

export default UploadFile;
