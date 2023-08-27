import "./dashboard.css";
import PlayerDashboard from "./playerDashboard";
import { useState, useEffect } from "react";

export default function Dashboard({ data }) {
  const [listClient, setListClient] = useState([]);
  useEffect(() => {
    if (data?.action === "join-room" || data?.action === "leave-room") {
      setListClient(data.payload.clients);
    }
  }, [data]);
  return (
    <div className="dashboard">
      <div className="dashboard-players">
        {listClient && listClient.length > 0 ? (
          listClient.map((item, index) => (
            <PlayerDashboard key={index} client={item} />
          ))
        ) : (
          <div className="no-player">No player</div>
        )}
      </div>
    </div>
  );
}
