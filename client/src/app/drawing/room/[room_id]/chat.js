import "./chat.css";
import { useState, useEffect } from "react";

export default function Chat({ data, room_id, ws, username, avatar_url }) {
  const [messagesChat, setMessagesChat] = useState([]);
  const [messagesGuess, setMessagesGuess] = useState([]);

  useEffect(() => {
    if (data?.action === "chat") {
      setMessagesChat((messagesChat) => [
        ...messagesChat,
        {
          username: data.sender.client_name,
          content: data.payload.message,
        },
      ]);
    } else if (data?.action === "join-room") {
      setMessagesChat((messagesChat) => [
        ...messagesChat,
        {
          username: "Room Notification",
          content: data.payload.message,
        },
      ]);
    } else if (data?.action === "leave-room") {
      setMessagesChat((messagesChat) => [
        ...messagesChat,
        {
          username: "Room Notification",
          content: data.payload.message,
        },
      ]);
    }
  }, [data]);

  const sendMessage = (event) => {
    if (event.key === "Enter" && event.target.value !== "") {
      if (event.target.placeholder === "Chat here") {
        ws.send({
          action: "chat",
          payload: event.target.value,
          sender: {
            client_name: username,
            avatar_url: avatar_url,
          },
          target: {
            room_id: room_id,
          },
        });
        event.target.value = "";
      } else if (event.target.placeholder === "Guess here") {
        event.target.value = "";
      }
    }
  };

  return (
    <div className="chat">
      <div className="left-chat chat-container">
        <div className="chat-messages">
          {messagesChat.map((message, index) => {
            return (
              <div key={index} className="chat-message">
                <div className="chat-username">{message.username} :</div>
                <div className="chat-content">{message.content}</div>
              </div>
            );
          })}
        </div>
        <input
          type="text"
          className="chat-input"
          placeholder="Chat here"
          onKeyDown={sendMessage}
        />
      </div>
      <div className="right-chat chat-container">
        <div className="chat-messages">
          {messagesGuess.map((message, index) => {
            return (
              <div key={index} className="chat-message">
                <div className="chat-username">{message.username} :</div>
                <div className="chat-content">{message.content}</div>
              </div>
            );
          })}
        </div>
        <input
          type="text"
          className="chat-input"
          placeholder="Guess here"
          onKeyDown={sendMessage}
        />
      </div>
    </div>
  );
}
