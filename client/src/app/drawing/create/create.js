import Link from "next/link";
import Image from "next/image";
import "./create.css";
// import DropdownComponent from "@/app/components/dropdown";
import Socket from "@/app/components/socket";

export default function CreateComponent() {
  const handleCreateRoom = () => {
    if (sessionStorage.getItem("username") === "") {
      sessionStorage.setItem(
        "username",
        `Alpha${Math.floor(Math.random() * 100000)}`
      );
    }
    const wsSocket = new Socket("room");
    wsSocket.open();
    wsSocket.send({
      action: "create-room",
      sender: { client_name: sessionStorage.getItem("username") },
      target: { room_id: "", max_players: 10, private: false },
      payload: "",
    });
    wsSocket.receive((data) => {
      console.log(data);
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
