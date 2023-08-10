export default function SideNavbar({ children }) {
    return (
        <div className="sidenavbar-container">
            <div className="sidenavbar">
                hi
                <nav>
                    <ul> Home </ul>
                    <ul> About </ul>
                    <ul> Contact </ul>
                </nav>
            </div>
            <div className="page-content">
                {children}
            </div>
        </div>
    )
}