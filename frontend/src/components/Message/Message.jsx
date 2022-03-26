// src/components/Message/Message.jsx
import React, { Component } from "react";
import "./Message.scss";

class Message extends Component {
  constructor(props) {
    super(props);
    let message_json = JSON.parse(this.props.message);
    let message_string = this.props.message;
    this.state = {
      message: message_json,
      stringData: message_string
    };
  }

  render() {
    return (
      <div className="Message">
        <div id="Message_ChatType">
          <label>Chat Type: {this.state.message.chatType}</label>
        </div>
        <div id="Message_ClientID">
          <label>Client ID: {this.state.message.clientID}</label>
        </div>
        <div id="Message_Body">
          <label>Message: {this.state.message.body}</label>
        </div>
        {/* {this.state.stringData} */}
      </div>
    );
  }
}

export default Message;