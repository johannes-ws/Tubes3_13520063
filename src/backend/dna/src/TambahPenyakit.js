import React, { useRef } from 'react'
import './GeneralStyle.css'
import { sendMsg } from "./Api";
export default function TambahPenyakit() {
  const namaPenyakit = useRef()
  const inputFilePenyakit = useRef()
  let inputText;

  const showFile = (e) => {
    e.preventDefault();
    const reader = new FileReader();
    reader.onload = (e) => {
      const text = e.target.result;
      console.log(text);
      inputText = text;
    };
    reader.readAsText(e.target.files[0]);
  };

  function handleSubmit(e) {
    sendMsg("Tambah"+ ";" + namaPenyakit.current.value+ ";" + inputText)
    namaPenyakit.current.value = null;
    inputFilePenyakit.current.value = null;
  };

  return (
    <>
      <h1 className='headera'>
        Tambah Penyakit
      </h1>
      <div>
         <h3 className='headerb'> Nama Penyakit : </h3> 
         <input className='pad' ref={namaPenyakit} type="text"></input>
      </div>
      <div>
        <h3 className='headerb'> Sequence DNA : </h3>
        <input className='pad' type='file' id='file' ref={inputFilePenyakit} onChange={showFile}/>
      </div>
      <div>
        <button className='button' onClick={handleSubmit}> Submit </button>
      </div>
    </>
  )
}
