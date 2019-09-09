import "./App.css";
import React from "react";
import Home from "./components/Home";
import Provider from "./components/Provider";

const App: React.FC = () => {
  return (
    <div className="App">
      <Provider>
        <Home />
      </Provider>
    </div>
  );
};

export default App;
