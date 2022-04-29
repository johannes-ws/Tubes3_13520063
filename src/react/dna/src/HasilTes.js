import React, { useRef } from 'react'
import ListQuery from './ListQuery'

export default function HasilTes() {
  const query = useRef()
  
  function handleSearch() {

  }
  return (
    <div>
      <h1> Hasil Tes </h1>
      <div>
        <input ref={query} type="text"></input>
        <button onClick={handleSearch}> Search </button>
      </div>
      <ListQuery query={query.current}/>
    </div>
  )
}
