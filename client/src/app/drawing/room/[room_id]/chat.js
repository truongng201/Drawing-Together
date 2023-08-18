import Socket from "../../../components/socket";
import "./chat.css";

import { useState, useEffect } from "react";

export default function Chat() {
  const [messagesChat, setMessagesChat] = useState([]);
  const [messagesGuess, setMessagesGuess] = useState([]);

  //setup websocket
  // const wsGuess = new Socket("messagesGuess");

  const wsChat = new Socket("messagesChat");
  useEffect(() => {
    wsChat.open();
    wsChat.receive((data) => {
      console.log(data);
      setMessagesChat([
        ...messagesChat,
        { username: data.username, content: data.message },
      ]);
    });
    return () => {
      wsChat.close();
    };
  }, [messagesChat]);

  const sendMessage = (event) => {
    if (event.key === "Enter" && event.target.value !== "") {
      if (event.target.placeholder === "Chat here") {
        wsChat.send(
          JSON.stringify({
            email: "truongng",
            username: "truongng",
            message: event.target.value,
          })
        );

        event.target.value = "";
      } else if (event.target.placeholder === "Guess here") {
        // wsGuess.sendMessage(event.target.value);
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
