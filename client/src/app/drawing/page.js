"use client";

import "./drawing.css";
import Login from "./login";
import Image from "next/image";

export default function Drawing() {
  return (
    <div className="drawing-container">
      <div className="drawing-layout">
        <div className="drawing-left-container">
          <div className="drawing-sub-container">
            <Login />
          </div>
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
