'use client';

import './room.css'
import Canvas2D from "./canvas2D";
import Chat from "./chat";
import Dashboard from "./dashboard";

export default function Room() {
    return (
        <div className="room-container">
            <Dashboard />
            <Canvas2D />
            <Chat />
        </div>
    )
}