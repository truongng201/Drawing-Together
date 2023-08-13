import './chat.css';

import { useState } from 'react';

export default function Chat() {
    const [messagesChat, setMessagesChat] = useState([]);
    const [messagesGuess, setMessagesGuess] = useState([]);

    const sendMessage = (event) => {

        if (event.key === 'Enter' && event.target.value !== '') {
            if (event.target.placeholder === 'Guess here') {
                setMessagesChat([...messagesChat, { username: 'test', content: event.target.value }]);
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