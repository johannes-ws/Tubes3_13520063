import React, { useRef } from 'react'
import './GeneralStyle.css'

export default function TambahPenyakit() {
  const namaPenyakit = useRef()
  const inputFilePenyakit = useRef()
  
  function handleSubmit(e) {
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
        <input className='pad' type='file' id='file' ref={inputFilePenyakit}/>
      </div>
      <div>
        <button className='button' onClick={handleSubmit}> Submit </button>
      </div>
    </>
  )
}
