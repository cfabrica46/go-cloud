let btnUpload = document.getElementById("upload");
btnUpload.addEventListener("click", uploadFile);

function uploadFile() {
    let inputFile = document.getElementById("input-file");
    let file = inputFile.files[0];

    fetch("/api/v1/upload", {
        method: "POST",
        body: JSON.stringify(file),
    })
        .then(() => {
            let title = document.getElementById("title");
            title.textContent = "Uploaded File";
        })
        .catch(() => {
            let title = document.getElementById("title");
            title.textContent = "Failure to Upload File";
        });
}
