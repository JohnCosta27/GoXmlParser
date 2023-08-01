const goWasm = new Go();

WebAssembly.instantiateStreaming(
  fetch("xml2json.wasm"),
  goWasm.importObject
).then((result) => {
  goWasm.run(result.instance);
  const json = XmlToJson("<hello><a>World</a><a>dsadsa</a></hello>");
  console.log(JSON.parse(json));
});
