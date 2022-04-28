import React, { useRef } from 'react'
import TambahPenyakit from './TambahPenyakit';
import TesDna from './TesDna';
import HasilTes from './HasilTes';

export default function MainMenu() {


    function handleTambahPenyakit() {
        
        return (
            <div>
                <TambahPenyakit />
            </div>
        )
    }

    function handleTesDna() {
        return (
            <div>
                <TesDna />
            </div>
        )
    }

    function handleHasilTes() {
        return (
            <div>
                <HasilTes />
            </div>
        )
    }
    return (
        <>
            <div>
                <button id='tombolMenu' onClick={handleTambahPenyakit}> Tambah Penyakit </button>
                <TambahPenyakit/>
                
            </div>
            <div>
                <button onClick={handleTesDna}> Tes DNA </button>
                <TesDna/>
            </div>
            <div>
                <button onClick={handleHasilTes}> Hasil Tes </button>
                <HasilTes/>
            </div>
        </>
  )
}
