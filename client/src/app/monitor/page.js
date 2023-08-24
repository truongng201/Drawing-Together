import "./monitor.css";

export default function Monitor() {
  return (
    <div className="monitor-container">
      <iframe
        src={process.env.NEXT_PUBLIC_MONITOR_DASHBOARD_SRC}
        title="monitor"
        width="100%"
        height="100%"
      />
    </div>
  );
}
