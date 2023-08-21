import Link from "next/link";
import "./login.css";
import Image from "next/image";
import { useEffect, useState } from "react";
import AlertComponent from "../components/alert";

export default function Login() {
  const [errorUsername, setErrorUsername] = useState(false);
  const checkUsername = (e) => {
    if (sessionStorage.getItem("username") === "") {
      e.preventDefault();
      setErrorUsername(true);
      setTimeout(() => {
        setErrorUsername(false);
      }, 5000);
    }
  };

  useEffect(() => {
    if (
      sessionStorage.getItem("username") === null ||
      sessionStorage.getItem("username") === ""
    ) {
      const username = `Alpha${Math.floor(Math.random() * 100000)}`;
      sessionStorage.setItem("username", username);
    }
  }, []);

  return (
    <div className="login-container">
      <div className="button-groups">
        <Link href="/drawing/join" onClick={checkUsername}>
          <div className="login-button join-button">
            <Image
              className="login-join-icon"
              alt="logo"
              src="/icons/door.png"
              width={30}
              height={30}
            />
            <span>Join</span>
          </div>
        </Link>

        <Link href="/drawing/create" onClick={checkUsername}>
          <div className="login-button create-button">
            <Image
              className="login-create-icon"
              alt="logo"
              src="/icons/game-controller.png"
              width={30}
              height={30}
            />
            <span>Create</span>
          </div>
        </Link>
      </div>
      <div>
        <div className="login-input">
          <div className="login-avatar-block">
            <Image
              className="login-avatar"
              alt="rand-avt"
              src="https://api.dicebear.com/6.x/shapes/svg?backgroundColor=b6e3f4,c0aede,121826"
              width={80}
              height={80}
            />
          </div>
          <div className="login-input-rand">
            <input
              type="text"
              placeholder="Username"
              defaultValue={sessionStorage.getItem("username")}
              onChange={(e) => {
                sessionStorage.setItem("username", e.target.value);
              }}
            />
          </div>
        </div>
        <div className="or-divider">
          <span className="line"></span>
          <span className="or-text">OR</span>
          <span className="line"></span>
        </div>
        <div className="oauth-groups">
          <div className="oauth-button google-button">
            <a href="google.com">
              <i className="fab fa-google"></i>
            </a>
          </div>
          <div className="oauth-button facebook-button">
            <a href="facebook.com">
              <i className="fab fa-facebook"></i>
            </a>
          </div>
          <div className="oauth-button github-button">
            <a href="github.com">
              <i className="fab fa-github"></i>
            </a>
          </div>
          <div className="oauth-button discord-button">
            <a href="discord.com">
              <i className="fab fa-discord"></i>
            </a>
          </div>
        </div>
      </div>
      {errorUsername && (
        <AlertComponent
          close={() => {
            setErrorUsername(false);
          }}
          message="Username is required"
        />
      )}
    </div>
  );
}
