import Link from "next/link";
import Image from "next/image";
import "./create.css";
// import DropdownComponent from "@/app/components/dropdown";
import Socket from "@/app/components/socket";
import { useRouter } from "next/navigation";

export default function CreateComponent() {
  const { push } = useRouter();

  const handleCreateRoom = () => {
    if (sessionStorage.getItem("username") === "") {
      sessionStorage.setItem(
        "username",
        `Alpha${Math.floor(Math.random() * 100000)}`
      );
    }
    const wsSocket = new Socket("room");
    wsSocket.open();

    setTimeout(() => {
      wsSocket.send(
        JSON.stringify({
          action: "create-room",
          payload: "create-room",
          sender: {
            client_name: sessionStorage.getItem("username"),
            client_id: "",
          },
          target: {
            room_id: "",
            max_players: 10,
            private: true,
          },
        })
      );
    }, 3000);
    wsSocket.receive((data) => {
      if (data.target?.room_id !== "" || data.target?.room_id !== undefined) {
        push(`/drawing/room/${data.target?.room_id}`);
      }
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
        <div className="create-button-page" onClick={handleCreateRoom}>
          Create
        </div>
      </div>
    </div>
  );
}
