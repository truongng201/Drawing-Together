import "./dashboard.css";
import PlayerDashboard from "./playerDashboard";

export default function Dashboard({ listClient }) {
  return (
    <div className="dashboard">
      <div className="dashboard-players">
        {listClient.length > 0 ? (
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
