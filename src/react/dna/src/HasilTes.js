import React, { useRef, useState } from 'react'
import ListQuery from './ListQuery'

export default function HasilTes() {
  const [hasil, setHasil] = useState([])
  const query = useRef()

  function handleSearch() {
    setHasil(['1','2'])
  }
  return (
    <div>
      <h1> Hasil Tes </h1>
      <div>
        <input ref={query} type="text"></input>
        <button onClick={handleSearch}> Search </button>
      </div>
      <ListQuery hasil={hasil}/>
    </div>
  )
}
