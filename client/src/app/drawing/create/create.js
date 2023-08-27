import Link from "next/link";
import Image from "next/image";
import "./create.css";
import { useState } from "react";
// import DropdownComponent from "@/app/components/dropdown";
import { useRouter } from "next/navigation";
import axios from "axios";
import AlertComponent from "@/app/components/alert";

export default function CreateComponent() {
  const { push } = useRouter();
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleCreateRoom = () => {
    if (sessionStorage.getItem("username") === "") {
      sessionStorage.setItem(
        "username",
        `Alpha${Math.floor(Math.random() * 100000)}`
      );
    }
    setIsLoading(true);

    axios
      .post(`${process.env.NEXT_PUBLIC_API_URI}/create-room`, {
        max_players: 10,
        private: false,
      })
      .then((res) => {
        const payload = res?.data;
        if (payload?.success) {
          if (payload?.data?.room_id) {
            push(`/drawing/room/${payload.data.room_id}`);
          }
        }
      })
      .catch((err) => {
        const payload = err?.response?.data;
        setError(payload?.message || "Something went wrong");
        setIsLoading(false);
      });
  };

  return (
    <div className="create-container">
      <div className="create-top">
        <Link href="/drawing">
          <Image
            src="/icons/back.png"
            width={30}
            height={30}
            className="back-icon"
            alt="back"
          />
        </Link>
        <div className="create-title">Let&apos;s create new room</div>
      </div>
      <div className="create-bottom">
        <div className="create-settings">
          {/* <DropdownComponent listItems={["public", "private"]} /> */}
        </div>
        {!isLoading ? (
          <div className="create-button-page" onClick={handleCreateRoom}>
            Create
          </div>
        ) : (
          <div className="create-button-page" aria-disabled>
            Loading...
          </div>
        )}
      </div>
      {error !== "" && (
        <AlertComponent message={error} close={() => setError("")} />
      )}
    </div>
  );
}
