import React from 'react'
import Query from './Query'

export default function ListQuery( {hasil} ) {
  return (
    hasil.map(element => {
      return <Query key={element} element={element}/>
    })
  )
}
