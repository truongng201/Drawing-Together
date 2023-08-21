"use client";

import "./room.css";
import Canvas2D from "./canvas2D";
import Chat from "./chat";
import Dashboard from "./dashboard";
import { useSearchParams } from "next/navigation";
import Socket from "@/app/components/socket";
import { useState } from "react";

export default function Room({ params }) {
  const searchParams = useSearchParams();
  const [listClient, setListClient] = useState([]); // [{client_name, client_id, avatar_url}

  const ws = new Socket("room");
  const username = sessionStorage.getItem("username");
  const avatar_url = `https://api.dicebear.com/6.x/bottts-neutral/svg?seed=${username}`;

  if (searchParams.get("action") === "create") {
    const ws = new Socket("room");
    ws.open();
    setTimeout(() => {
      ws.send({
        action: "create-room",
        payload: "create-room",
        sender: {
          client_name: username,
          client_id: "",
          avatar_url: avatar_url,
        },
        target: {
          room_id: params.room_id,
          max_players: 10,
          private: false,
        },
      });
    }, 3000);
    ws.receive((data) => {
      console.log(data);
    });
  } else if (searchParams.get("action") === "join") {
    ws.open();
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
          room_id: params.room_id,
          max_players: 10,
          private: false,
        },
      });
    }, 3000);
    ws.receive((data) => {
      console.log(data);
    });
  }

  return (
    <div className="room-container">
      <Dashboard listClient={listClient} />
      <Canvas2D />
      <Chat />
    </div>
  );
}
