import "./login.css";
import Image from "next/image";

export default function Login() {
    return (
        <div className="login-container">
            <div className="login-left-container">
                <Image
                    className="login-avatar"
                    alt="rand-avt"
                    src="https://api.dicebear.com/6.x/bottts-neutral/svg?seed=Loki"
                    width={80}
                    height={80}
                />
                <div className="login-input">
                    <Image
                        className="login-username-icon"
                        alt="rand-avt"
                        src="/icons/username.png"
                        width={20}
                        height={20}
                    />
                    <div className="login-username-text">Nickname</div>
                    <input type="text" placeholder="Username" />
                </div>
                <div className="button-groups">
                    <div className="login-button join-button">Join</div>
                    <div className="login-button create-button">Create</div>
                </div>
            </div>
            <div className="login-divider-vertical"></div>
            <div className="login-right-container">left right</div>
        </div>
    );
}
