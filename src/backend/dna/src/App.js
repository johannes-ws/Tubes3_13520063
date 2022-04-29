import React from 'react';
import MainMenu from './MainMenu';
import { connect } from "./Api";

function App() {
  connect();


  return (
    <div>
      <MainMenu />
    </div>
  );
}

export default App;
