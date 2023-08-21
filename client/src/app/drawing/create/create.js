import Link from "next/link";
import Image from "next/image";
import "./create.css";
// import DropdownComponent from "@/app/components/dropdown";
import { useRouter } from "next/navigation";
import { customAlphabet } from "nanoid";

export default function CreateComponent() {
  const { push } = useRouter();

  const handleCreateRoom = () => {
    if (sessionStorage.getItem("username") === "") {
      sessionStorage.setItem(
        "username",
        `Alpha${Math.floor(Math.random() * 100000)}`
      );
    }
    const nanoid = customAlphabet("1234567890abcdefghijklmnopqrstuvwxyz", 10);
    const room_id = nanoid(10);
    push(`/drawing/room/${room_id}?action=create`);
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
