"use client";

import "./room.css";
import Canvas2D from "./canvas2D";
import Chat from "./chat";
import Dashboard from "./dashboard";
import Socket from "@/app/components/socket";
import { useEffect, useState } from "react";

export default function Room({ params }) {
  const [listClient, setListClient] = useState([]); // [{client_name, client_id, avatar_url}

  const ws = new Socket("room");
  ws.open();
  const username = sessionStorage.getItem("username");
  const avatar_url = `https://api.dicebear.com/6.x/bottts-neutral/svg?seed=${username}`;
  const room_id = params.room_id;

  useEffect(() => {
    
   setTimeout(() => {
      ws.send({
        action: "join-room",
        payload: "join-room",
        sender: {
          client_name: username,
          client_id: "",
          avatar_url: avatar_url,
        },
        target: {
          room_id: room_id,
          max_players: 10,
          private: false,
        },
      });
    }, 1000);
  }, [username, avatar_url, room_id, ws]);

  return (
    <div className="room-container">
      <Dashboard listClient={listClient} />
      <Canvas2D />
      <Chat ws={ws} room_id={room_id} />
    </div>
  );
}
