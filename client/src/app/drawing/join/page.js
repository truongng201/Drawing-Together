"use client";

import "../drawing.css";
import Image from "next/image";
import JoinComponent from "./join";

export default function DrawingJoin() {
  return (
    <div className="drawing-container">
      <div className="drawing-layout">
        <div className="drawing-left-container">
          <JoinComponent />
        </div>
        <div className="drawing-right-container">
          <Image
            className="login-logo"
            alt="logo"
            src="/logo/logo-no-background.svg"
            width={300}
            height={100}
          />
        </div>
      </div>
    </div>
  );
}
