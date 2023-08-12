import './SideNavbar.css'
import Link from 'next/link'
import Image from 'next/image'

export default function SideNavbar() {
    const pages = [
        {
            name: 'Home',
            link: '/',
            className: 'page-component home',
            iconPath: '/icons/home.png'
        },
        {
            name: 'Drawing',
            link: '/drawing',
            className: 'page-component drawing',
            iconPath: '/icons/drawing.png'
        },
        {
            name: 'Monitor',
            link: '/monitor',
            className: 'page-component monitor',
            iconPath: '/icons/monitor.png'
        },
        {
            name: 'Blog',
            link: '/blog',
            className: 'page-component blog',
            iconPath: '/icons/blog.png'
        },
        {
            name: 'Settings',
            link: '/settings',
            className: 'page-component settings',
            iconPath: '/icons/settings.png'
        }
    ]
    return (
        <div className="sidenavbar">
            <Link className='logo' href='/'><img src="/logo/logo-no-background.svg" /></Link>
            <div className="pages">
                {pages.map(page => {
                    return (
                        <Link className={page.className} href={page.link}>
                            <Image width={16} height={16} src={page.iconPath} alt='page icon' />
                            <div>{page.name}</div>
                        </Link>
                    )
                })}
            </div>
            <div className='developer-info'>
                <div className='username'>Truong Nguyen</div>
                <Image
                    src="/avatar.png"
                    className='avatar'
                    alt='developer avatar'
                    width={100}
                    height={100}
                />
                <div className='bio'>
                    I&apos;m a software engineer who loves to build things. I write about building scalable systems.
                </div>
                <div className='social-media'>
                    <div className='github'>
                        <a href='https://github.com/truongng201'>
                            <i className="fab fa-github"></i>
                        </a>
                    </div>
                    <div className='facebook'>
                        <a href='https://facebook.com'>
                            <i className="fab fa-facebook"></i>
                        </a>
                    </div>
                    <div className='linkeding'>
                        <a href='https://linkedin.com'>
                            <i className="fab fa-linkedin"></i>
                        </a>
                    </div>
                </div>
            </div>
        </div>
    )
}