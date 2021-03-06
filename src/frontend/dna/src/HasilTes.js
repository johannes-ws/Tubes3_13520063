import React, { useRef, useState } from 'react'
import ListQuery from './ListQuery'
import './GeneralStyle.css'
import { sendMsg, message } from "./Api";

export default function HasilTes() {
  const [hasil, setHasil] = useState([])
  const query = useRef()

  function handleSearch() {
    sendMsg("Hasil" + ";" + query.current.value)
    console.log(message === null)
    console.log(message)
    if (message !== '' && message !== null) {
      
    }
  }
  return (
    <div>
      <h1 className='headera'> Hasil Tes </h1>
      <div>
        <input className='pad' ref={query} type="text"></input>
        <button onClick={handleSearch}> Search </button>
      </div>
      <ListQuery hasil={hasil}/>
    </div>
  )
}
