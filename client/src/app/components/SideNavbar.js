import './SideNavbar.css'
import Link from 'next/link'

export default function SideNavbar() {
    // create a list of pages to display
    const pages = [
        {
            name: 'Home',
            link: '/',
            className: 'home',
            icon: 'home'
        },
        {
            name: 'Drawing',
            link: '/drawing',
            className: 'drawing'
        },
        {
            name: 'Monitor',
            link: '/monitor',
            className: 'monitor'
        },
        {
            name: 'Blog',
            link: '/blog'
        }
    ]
    return (
        <div className="sidenavbar">
            <div>Logo</div>
            <div className="pages">
                {pages.map(page => {
                    return (
                        <div className={page.className}>
                            <Link href={page.link}>{page.name}</Link>
                        </div>
                    )
                })}
            </div>
            <div>
                <div>avatar</div>
                <div>Some quote here</div>
                <div>Github icon</div>
                <div>Facebook icon</div>
                <div>Linkedin icon</div>
            </div>
            <div>Settings icon</div>
        </div>
    )
}