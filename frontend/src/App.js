// App.js
import React, { Component } from "react";
import "./App.css";

// Components
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory />
        <ChatInput />
      </div>
    );
  }
}

export default App;