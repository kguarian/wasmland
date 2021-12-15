function taskCardAPI(mode, id, name, info) {
    switch (mode) {
        case 0: //add
            return GOAPI_addTaskCard(id, name, info);
        case 1:
            return GOAPI_deleteTaskCard(id);
        case 2:
            return GOAPI_setTaskCard(id, name);
        default:
            return -1;
    }
}

function initiateGo(path) {
    let wasmInstance;
    let wasmModule;
    if (!WebAssembly.instantiateStreaming) {
        WebAssembly.instantiateStreaming = async(resp, importObject) => {
            const source = await (await resp).arrayBuffer();
            return await WebAssembly.instantiate(source, importObject);
        };
    }
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch(path), go.importObject).then((result) => {
        wasmModule = result.module;
        wasmInstance = result.instance;

        go.run(wasmInstance);
    }).catch((err) => {
        console.error(err);
    });
    console.log("Go Webassembly library running.");
}

async function startTimer() {
    var id;
    var time;

    id = "timer_display";
    time = document.getElementById("timer_display").value;
    wasm_startTimer(id, time);
}

function stopTimer() {
    var id;
    id = "timer_display";
    wasm_stopTimer(id);
}

function resetTimer() {
    var id;
    var retVal;
    id = "timer_display";
    retVal = wasm_resetTimer(id);
}

class taskCard {
    constructor(desiredId, name, info) {
        this.name = name;
        this.desiredIdthis.info = info;

    }
};