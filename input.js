import "wasm_exec.js"

var base64String = "PLACE_BASE64_STRING_FROM_wasmstr.txt_HERE";

async function runWasm() {
    try {
        const vaultIndex = document.getElementById("vaultIndex").value;
        const chainInt = document.getElementById("chain").value;
        const subIndex = document.getElementById("subIndex").value;

        const wasmBinary = Uint8Array.from(atob(base64String), c => c.charCodeAt(0));
        const { instance } = await WebAssembly.instantiate(wasmBinary);

        const output = instance.exports.generateChildExtendedPrivateKey("./metadata.json", 0, vaultIndex, chainInt, subIndex);
        document.getElementById("output").innerText = "Child Extended Private Key: " + output;
    } catch (error) {
        console.error('An error occurred:', error);
        // Handle the error, e.g., display an error message to the user
    }
}

// Example: Add event listener to call runWasm function on button click
//document.getElementById('runButton').addEventListener('click', runWasm);
