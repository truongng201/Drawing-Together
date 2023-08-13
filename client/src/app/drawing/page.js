'use client'

import Canvas2D from "./canvas2D";
import Chat from "./chat";
import Dashboard from "./dashboard";
import './drawing.css';

export default function Drawing() {
    return (
        <div className="drawing-container">
            <Dashboard />
            <Canvas2D />
            <Chat />
        </div>
    )
}