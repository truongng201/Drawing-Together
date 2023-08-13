import Socket from '../components/socket';
import './chat.css';

import { useState, useEffect } from 'react';

export default function Chat() {
    const [messagesChat, setMessagesChat] = useState([]);
    const [messagesGuess, setMessagesGuess] = useState([]);

    //setup websocket
    const ws = new Socket();

    useEffect(() => {
        ws.openSocket();
        ws.reconnectSocket();
        ws.receiveMessage(
            (data) => {
                setMessagesChat([...messagesChat, { username: 'test2', content: data }]);
            }
        );
        return () => {
            ws.closeSocket();
        }
    }, [messagesChat]);

    const sendMessage = (event) => {
        if (event.key === 'Enter' && event.target.value !== '') {
            if (event.target.placeholder === 'Guess here') {
                setMessagesChat([...messagesChat, { username: 'test1', content: event.target.value }]);
                ws.sendMessage(event.target.value);
                event.target.value = '';
            }
            else if (event.target.placeholder === 'Chat here') {
                setMessagesGuess([...messagesGuess, { username: 'test', content: event.target.value }]);
                event.target.value = '';
            }
        }
    }

    return (
        <div className="chat">
            <div className='left-chat chat-container'>
                <div className='chat-messages'>
                    {messagesChat.map((message, index) => {
                        return (
                            <div key={index} className='chat-message'>
                                <div className='chat-username'>{message.username} :</div>
                                <div className='chat-content'>{message.content}</div>
                            </div>
                        )
                    })}
                </div>
                <input
                    type="text"
                    className="chat-input"
                    placeholder='Guess here'
                    onKeyDown={sendMessage}
                />
            </div>
            <div className='right-chat chat-container'>
                <div className='chat-messages'>
                    {messagesGuess.map((message, index) => {
                        return (
                            <div key={index} className='chat-message'>
                                <div className='chat-username'>{message.username} :</div>
                                <div className='chat-content'>{message.content}</div>
                            </div>
                        )
                    })}
                </div>
                <input
                    type="text"
                    className="chat-input"
                    placeholder='Chat here'
                    onKeyDown={sendMessage}
                />
            </div>
        </div>
    )
}