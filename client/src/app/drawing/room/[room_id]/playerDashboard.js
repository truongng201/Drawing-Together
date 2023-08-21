import Image from "next/image";

export default function PlayerDashboard({ client }) {
  return (
    <div className="player-dashboard">
      <Image
        className="player-avatar"
        src={client?.avatar_url}
        width={50}
        height={50}
        alt="user-avatar"
      />
      <div className="player-name">{client?.client_name}</div>
    </div>
  );
}
