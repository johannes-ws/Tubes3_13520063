import React, { useRef } from 'react'

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
        <h1> Tes DNA </h1>
        <div>
          Nama Pengguna :
          <input ref={namaPengguna} type="text"></input>
        </div>
        <div>
        Sequence DNA :
        <input type='file' id='file' ref={inputFilePasien}/>
        </div>
        <div>
          Prediksi Penyakit :
          <input ref={prediksiPenyakit} type="text"></input>
        </div>
        <div>
          <button onClick={handleSubmit}> Submit </button>
        </div>
      </div>
  )
}
