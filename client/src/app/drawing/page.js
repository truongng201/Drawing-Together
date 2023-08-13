'use client'

import Canvas2D from "./canvas2D";

export default function Drawing() {
    return (
        <div className="drawing-container">

            <div className="drawing">
                <div className="drawing-header">
                    <h1>Drawing</h1>
                </div>
                <Canvas2D />
            </div>
        </div>
    )
}