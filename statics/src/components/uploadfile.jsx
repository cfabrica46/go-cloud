import React, { useState } from "react";
const UploadFile = () => {
    const [selectedFile, setSelectedFile] = useState(undefined);

    const handleSubmit = () => {
        console.log(selectedFile);

        let dataForm = new FormData();
        dataForm.append("file", selectedFile);

        fetch("http://localhost:8080/api/v1/upload", {
            method: "POST",
            body: dataForm,
        })
            .then((responsive) => {
                if (responsive.status >= 400) {
                    throw true;
                }
            })
            .then(() => {
                let title = document.getElementById("title");
                title.textContent = "Uploaded File";
            })
            .catch(() => {
                let title = document.getElementById("title");
                title.textContent = "Failure to Upload File";
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
