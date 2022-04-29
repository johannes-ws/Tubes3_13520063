import React, { useRef } from 'react'
import './GeneralStyle.css'

export default function TesDna() {
  const namaPengguna = useRef()
  const inputFilePasien = useRef()
  const prediksiPenyakit = useRef()

  function handleSubmit(e) {
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
        <input className='pad' type='file' id='file' ref={inputFilePasien}/>
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
