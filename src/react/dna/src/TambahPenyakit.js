import React, { useRef } from 'react'

export default function TambahPenyakit() {
  const namaPenyakit = useRef()
  const inputFilePenyakit = useRef()
  
  function handleSubmit(e) {
    namaPenyakit.current.value = null;
    inputFilePenyakit.current.value = null;
  };

  return (
    <>
      <h1>
        Tambah Penyakit
      </h1>
      <div>
        Nama Penyakit :
        <input ref={namaPenyakit} type="text"></input>
      </div>
      <div>
        Sequence DNA :
        <input type='file' id='file' ref={inputFilePenyakit}/>
      </div>
      <div>
        <button onClick={handleSubmit}> Submit </button>
      </div>
    </>
  )
}
