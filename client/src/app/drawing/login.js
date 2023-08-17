import "./login.css";
import Image from "next/image";

export default function Login() {
  return (
    <div className="login-container">
      <div className="login-left-container">
        <div className="button-groups">
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
              <input type="text" placeholder="Username" />
            </div>
          </div>
          <div className="or-divider">
            <span className="line"></span>
            <span className="or-text">OR</span>
            <span className="line"></span>
          </div>
          <div className="oauth-groups">
            <div className="oauth-button google-button">Google</div>
            <div className="oauth-button facebook-button">Facebook</div>
            <div className="oauth-button github-button">Github</div>
            <div className="oauth-button discord-button">Discord</div>
          </div>
        </div>
      </div>
      <div className="login-right-container">
        <Image
          className="login-logo"
          alt="logo"
          src="/logo/logo-no-background.svg"
          width={300}
          height={100}
        />
      </div>
    </div>
  );
}
