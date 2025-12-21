async function rest(url) {
    const res = await fetch(url)
    return await res.json();
}

async function getConfig() {
    return await rest("data/config.json");
}

async function onSubmit() {
    const nameElement = document.getElementById("name");
    const messageElement = document.getElementById("message");
    try {
        const config = await getConfig();
        const nameValue = nameElement.value || "";
        const res = await rest(`${config.api}/greet?name=${nameValue}`);
        messageElement.innerHTML = res.message;
    } catch (e) {
        messageElement.innerHTML = e.message;
    }
}
