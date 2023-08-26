"use client";

import "./room.css";
import Canvas2D from "./canvas2D";
import Chat from "./chat";
import Dashboard from "./dashboard";
import Socket from "@/app/components/socket";
import { useEffect, useState, useRef } from "react";

export default function Room({ params }) {
  const { current: instance } = useRef({});
  const ws = (instance.ws = instance.ws || new Socket("room"));
  const [wsData, setWsData] = useState(null); // [{client_name, client_id, avatar_url}

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
          avatar_url: avatar_url,
        },
        target: {
          room_id: room_id,
        },
      });
    }, 1000);
    ws.receive((data) => {
      setWsData(data);
    });
  }, [username, avatar_url, room_id, ws]);

  return (
    <div className="room-container ">
      {wsData ? (
        <>
          <Dashboard listClient={wsData?.payload.clients || []} />
          <Canvas2D />
          <Chat data={wsData} ws={ws} room_id={room_id} />
        </>
      ) : (
        <div>Loading...</div>
      )}
    </div>
  );
}
