import React, { useState } from 'react'
import TambahPenyakit from './TambahPenyakit';
import TesDna from './TesDna';
import HasilTes from './HasilTes';

export default function MainMenu() {
    const [viewMode, setViewMode] = useState(0)

    function handleTambahPenyakit() {
        setViewMode(1);
    }

    function handleTesDna() {
        setViewMode(2);
    }

    function handleHasilTes() {
        setViewMode(3);
    }

    function handleBack() {
        setViewMode(0);
    }


    if (viewMode === 0) {
        return (
            <>
                <h1> Main Menu </h1>
                <div>
                    <button onClick={handleTambahPenyakit}> Tambah Penyakit </button>
                </div>
                <div>
                    <button onClick={handleTesDna}> Tes DNA </button>
                </div>
                <div>
                    <button onClick={handleHasilTes}> Hasil Tes </button>
                </div>
            </>
        )
    }

    if (viewMode === 1) {
        return (
            <>
                <div>
                    <button onClick={handleBack}> Back </button>
                </div>
                <TambahPenyakit/>
            </>
        )
    }

    if (viewMode === 2) {
        return (
            <>
                <div>
                    <button onClick={handleBack}> Back </button>
                </div>
                <TesDna/>
            </>
        )
    }

    if (viewMode === 3) {
        return (
            <>
                <div>
                    <button onClick={handleBack}> Back </button>
                </div>
                <HasilTes/>
            </>
        )
    }
}
