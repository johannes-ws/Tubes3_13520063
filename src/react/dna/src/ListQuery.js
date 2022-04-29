import React from 'react'
import Query from './Query'

export default function ListQuery( {hasil} ) {
  return (
    <ul>
      {
        hasil.map(element => {
          return <Query key={element} element={element}/>
        })
      }
    </ul>
  )
}
