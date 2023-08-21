"use client";

import "./room.css";
import Canvas2D from "./canvas2D";
import Chat from "./chat";
import Dashboard from "./dashboard";
import { useSearchParams, usePathname } from "next/navigation";

export default function Room({ params }) {
  const searchParams = useSearchParams();
  console.log("params", params.room_id, searchParams.get("action"));

  //   if (params.get("action") === "create") {
  //   }

  return (
    <div className="room-container">
      <Dashboard />
      <Canvas2D />
      <Chat />
    </div>
  );
}
