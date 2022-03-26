import React, { Component } from "react";
import "./ChatHistory.scss";
import Message from '../Message/Message';
import { connect } from '../../api';

class ChatHistory extends Component {
    constructor(props) {
        super(props);
        this.state = {
            chatHistory: [],
        }
    }

    componentDidMount() {
        connect((msg) => {
            this.setState({
            chatHistory: [...this.state.chatHistory, msg]
            })
        });
    }

    render() {
        const messages = this.state.chatHistory.map((msg, index) => <Message key={index} message={msg.data} />);

        return (
            <div className='ChatHistory'>
                <h2>Chat History</h2>
                {messages}
            </div>
        );
    };
}

export default ChatHistory;