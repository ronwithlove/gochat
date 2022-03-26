import React, { Component } from "react";
import { sendMsg } from '../../api';
import "./ChatInput.scss";

class ChatInput extends Component {
  constructor(props) {
    super(props)
    this.state = {
      chatType: 'Private',
      clientID: '',
      message:  '',
    };

    this.handleSendMessage = this.handleSendMessage.bind(this);
    this.handleSendWithMouse = this.handleSendWithMouse.bind(this);
    this.handleSendWithKeyboard = this.handleSendWithKeyboard.bind(this);
  }

  handleSendMessage(type, id, msg) {
    var message = {
      chatType: type,
      clientId: id,
      message: msg
    };
    console.log(message);
    sendMsg(JSON.stringify(message));
  }

  handleSendWithKeyboard(event) {
    if(event.keyCode === 13) {
      if (this.state.clientID === '' || this.state.message === '') {
        alert('Please fill out both input fields before sending the message.');
      } else {
        this.handleSendMessage(this.state.chatType, this.state.clientID, this.state.message);
        this.state.clientID = "";
        this.state.message = "";
        this.forceUpdate();
      }
    }
  }

  handleSendWithMouse() {
    if (this.state.clientID === '' || this.state.message === '') {
      alert('Please fill out both input fields before sending the message.');
    } else {
      this.handleSendMessage(this.state.chatType, this.state.clientID, this.state.message);
      this.state.clientID = "";
      this.state.message = "";
      this.forceUpdate();
    }
  }

  render() {
    return (
      <div className="ChatInput">
        <div id="ChatInput_Item">
          <div id="ChatInput_Item_Title">
            <label>Chat Type:</label>
          </div>
          <select 
            onChange={event => this.setState({chatType: event.target.value})}
          >
            <option>Private</option>
            <option>Public</option>
          </select>
        </div>
        <div id="ChatInput_Item">
          <div id="ChatInput_Item_Title">
            <label>Client ID:</label>
          </div>
          <input 
            value={this.state.clientID} 
            onChange={event => this.setState({clientID: event.target.value})} 
            onKeyDown={this.handleSendWithKeyboard}
          />
        </div>
        <div id="ChatInput_Item">
          <div id="ChatInput_Item_Title">
            <label>Message:</label>
          </div>
          <input 
            value={this.state.message} 
            onChange={event => this.setState({message: event.target.value})} 
            onKeyDown={this.handleSendWithKeyboard}
          />
        </div>
        <div id="ChatInput_Button">
          <button onClick={this.handleSendWithMouse}>SEND</button>
        </div>
      </div>
    );
  }
}

export default ChatInput;