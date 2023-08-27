import { useEffect, useRef, useState } from "react";
import "./canvas2D.css";

export default function Canvas2D({ ws, room_id, username, avatar_url, data }) {
  const canvasRef = useRef(null);
  const [drawing, setDrawing] = useState(false);
  const [context, setContext] = useState(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    const ctx = canvas.getContext("2d");
    canvas.style.width = "98%";
    canvas.style.height = "95%";
    canvas.width = canvas.offsetWidth;
    canvas.height = canvas.offsetHeight;
    console.log(data);
    if (data?.action === "drawing") {
      console.log(data);
      if (data.payload.state === "start") {
        context.beginPath();
        context.moveTo(data.payload.offset_x, data.payload.offset_y);
      } else if (data.payload.state === "continue") {
        context.lineTo(data.payload.offset_x, data.payload.offset_y);
        context.stroke();
      } else if (data.payload.state === "stop") {
        context.closePath();
      }
    }

    setContext(ctx);
  }, [data]);

  const startDrawing = (event) => {
    const { offsetX, offsetY } = event.nativeEvent;
    context.beginPath();
    context.moveTo(offsetX, offsetY);
    ws.send({
      action: "drawing",
      payload: {
        offset_x: offsetX,
        offset_y: offsetY,
        state: "start",
      },
      sender: {
        client_name: username,
        avatar_url: avatar_url,
      },
      target: {
        room_id: room_id,
      },
    });
    setDrawing(true);
  };

  const continueDrawing = (event) => {
    if (!drawing) return;

    const { offsetX, offsetY } = event.nativeEvent;
    context.lineTo(offsetX, offsetY);
    context.stroke();
    ws.send({
      action: "drawing",
      payload: {
        offset_x: offsetX,
        offset_y: offsetY,
        state: "continue",
      },
      sender: {
        client_name: username,
        avatar_url: avatar_url,
      },
      target: {
        room_id: room_id,
      },
    });
  };

  const stopDrawing = () => {
    context.closePath();
    ws.send({
      action: "drawing",
      payload: {
        offset_x: 0,
        offset_y: 0,
        state: "stop",
      },
      sender: {
        client_name: username,
        avatar_url: avatar_url,
      },
      target: {
        room_id: room_id,
      },
    });
    setDrawing(false);
  };
  return (
    <canvas
      ref={canvasRef}
      className="canvas2D"
      onMouseDown={startDrawing}
      onMouseMove={continueDrawing}
      onMouseUp={stopDrawing}
      onMouseOut={stopDrawing}
    />
  );
}
