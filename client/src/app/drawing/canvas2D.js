import { useEffect, useRef, useState } from 'react';
import './canvas2D.css';

export default function Canvas2D(){
    const canvasRef = useRef(null);
    const [drawing, setDrawing] = useState(false);
    const [context, setContext] = useState(null);

    useEffect(() => {
        const canvas = canvasRef.current;
        const ctx = canvas.getContext('2d');
        canvas.style.width = "98%";
        canvas.style.height = "95%";
        canvas.width = canvas.offsetWidth;
        canvas.height = canvas.offsetHeight;
        setContext(ctx);
    }, []);

    const startDrawing = (event) => {
        const { offsetX, offsetY } = event.nativeEvent;
        context.beginPath();
        context.moveTo(offsetX, offsetY);
        setDrawing(true);
    };

    const continueDrawing = (event) => {
        if (!drawing) return;

        const { offsetX, offsetY } = event.nativeEvent;
        context.lineTo(offsetX, offsetY);
        context.stroke();
    };

    const stopDrawing = () => {
        context.closePath();
        setDrawing(false);
    };
    return (
        <canvas
            ref={canvasRef}
            className='canvas2D'
            onMouseDown={startDrawing}
            onMouseMove={continueDrawing}
            onMouseUp={stopDrawing}
            onMouseOut={stopDrawing}
        />
    )
}