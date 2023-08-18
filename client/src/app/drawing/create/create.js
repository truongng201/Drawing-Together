import Link from "next/link";
import Image from "next/image";
import "./create.css";
import DropdownComponent from "@/app/components/dropdown";

export default function CreateComponent() {
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
        <div className="create-button-page">Create</div>
      </div>
    </div>
  );
}
