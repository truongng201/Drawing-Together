import Link from "next/link";
import Image from "next/image";
import "./join.css";
import { useState } from "react";
import { useRouter } from "next/navigation";
import AlertComponent from "@/app/components/alert";
import axios from "axios";

export default function JoinComponent() {
  const [roomId, setRoomId] = useState("");
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const { push } = useRouter();
  const handleJoinRoom = (event) => {
    if (event.key === "Enter" || event.type === "click") {
      setIsLoading(true);
      if (roomId === "") {
        setError("Please enter room id");
        setIsLoading(false);
        setTimeout(() => {
          setError(false);
        }, 10000);
      } else {
        axios
          .post(`${process.env.NEXT_PUBLIC_API_URI}/check-room-existed`, {
            room_id: roomId,
          })
          .then((res) => {
            const payload = res?.data;
            if (payload?.success) {
              push(`/drawing/room/${roomId}`);
            }
          })
          .catch((err) => {
            const payload = err?.response?.data;
            setError(payload?.message || "Something went wrong");
            setIsLoading(false);
          });
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
            onKeyDown={handleJoinRoom}
            onChange={(e) => {
              setRoomId(e.target.value);
            }}
          />
        </div>
        {!isLoading ? (
          <div className="join-button-page" onClick={handleJoinRoom}>
            Join
          </div>
        ) : (
          <div className="join-button-page" aria-disabled>
            Loading...
          </div>
        )}
      </div>
      {error !== "" && (
        <AlertComponent
          message={error}
          close={() => {
            setError("");
          }}
        />
      )}
    </div>
  );
}
