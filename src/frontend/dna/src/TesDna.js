import React, { useRef } from 'react'
import './GeneralStyle.css'
import { sendMsg } from "./Api";
export default function TesDna() {
  const namaPengguna = useRef()
  const inputFilePasien = useRef()
  const prediksiPenyakit = useRef()
  let inputSequence;

  const showFile = (e) => {
    e.preventDefault();
    const reader = new FileReader();
    reader.onload = (e) => {
      const text = e.target.result;
      console.log(text);
      inputSequence = text;
    };
    reader.readAsText(e.target.files[0]);
  };

  function handleSubmit(e) {
    sendMsg("Test" + ";" + namaPengguna.current.value + ";" + inputSequence + ";" + prediksiPenyakit.current.value)
    namaPengguna.current.value = null;
    inputFilePasien.current.value = null;
    prediksiPenyakit.current.value = null;
  };
  return (
      <div>
        <h1 className='headera'> Tes DNA </h1>
        <div>
          <h3 className='headerb'>Nama Pengguna :</h3>
          <input className='pad' ref={namaPengguna} type="text"></input>
        </div>
        <div>
        <h3 className='headerb'>Sequence DNA : </h3>
        <input className='pad' type='file' id='file' ref={inputFilePasien} onChange={showFile}/>
        </div>
        <div>
          <h3 className='headerb'>Prediksi Penyakit : </h3>
          <input className='pad' ref={prediksiPenyakit} type="text"></input>
        </div>
        <div>
          <button className='button' onClick={handleSubmit}> Submit </button>
        </div>
      </div>
  )
}
