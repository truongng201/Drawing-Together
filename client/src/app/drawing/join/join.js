import Link from "next/link";
import Image from "next/image";
import "./join.css";
import { useState } from "react";
import { useRouter } from "next/navigation";
import AlertComponent from "@/app/components/alert";

export default function JoinComponent() {
  const [roomId, setRoomId] = useState("");
  const [errorRoomId, setErrorRoomId] = useState(false);
  const { push } = useRouter();
  const joinRoom = (event) => {
    if (event.key === "Enter" || event.type === "click") {
      if (roomId === "") {
        setErrorRoomId(true);
        setTimeout(() => {
          setErrorRoomId(false);
        }, 10000);
      } else {
        push(`/drawing/room/${roomId}?action=join`);
      }
    }
  };

  return (
    <div className="join-container">
      <div className="join-top">
        <Link href="/drawing">
          <Image
            src="/icons/back.png"
            width={30}
            height={30}
            className="back-icon"
            alt="back"
          />
        </Link>
        <div className="join-title">Let&apos;s join a room</div>
      </div>

      <div className="join-bottom">
        <div className="join-input">
          <input
            type="text"
            placeholder="Room ID"
            onKeyDown={joinRoom}
            onChange={(e) => {
              setRoomId(e.target.value);
            }}
          />
        </div>
        <div className="join-button-page" onClick={joinRoom}>
          Join
        </div>
      </div>
      {errorRoomId && (
        <AlertComponent
          message="Please enter room id"
          close={() => {
            setErrorRoomId(false);
          }}
        />
      )}
    </div>
  );
}
