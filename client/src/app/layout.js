import SideNavbar from './components/SideNavbar'
import './globals.css'
import { Inter } from 'next/font/google'
import Script from 'next/script'

const inter = Inter({ subsets: ['latin'] })

export const metadata = {
  title: 'Alpha',
}

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div className='left-container'>
          <SideNavbar />
        </div>
        <div className='right-container flex min-h-screen flex-col'>{children}</div>
      </body>
      <Script id='script-width'>
        {`window.innerWidth < 1200 && window.location.replace('/404')`}
      </Script>
    </html>
  )
}
