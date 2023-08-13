import './chat.css';

export default function Chat() {
    return (
        <div className="chat">
            <div className='left-chat chat-container'>
                <div className='chat-messages'>
                    Text


                </div>
                <input type="text" className="chat-input" placeholder='Guess here' />
            </div>
            <div className='right-chat chat-container'>
                <div className='chat-messages'>
                    Text
                </div>
                <input type="text" className="chat-input" placeholder='Chat here' />
            </div>
        </div>
    )
}