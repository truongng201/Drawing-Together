export default function Monitor() {
    return (
        <div className="monitor-container">

            <div className="monitor">
                <div className="monitor-header">
                    {
                        [...Array(100)].map((e, i) => {
                            return (
                                <div className="monitor-header-item">
                                    <h1>monitor</h1>
                                </div>
                            )
                        })
                    }
                </div>
            </div>
        </div>
    )
}    