import { NavLink, useNavigate } from 'react-router-dom'
import Button from '../button/button'
import styles from './sidebar.module.css'

function navItemClassName({ isActive }: { isActive: boolean }): string {
  return `${styles.navItem} ${isActive ? styles.navItemActive : ''}`
}

export default function Sidebar() {
  const navigate = useNavigate()

  function handleLogout() {
    localStorage.removeItem('token')
    navigate('/login', { replace: true })
  }

  return (
    <aside className={styles.sidebar}>
      <div>
        <p className={styles.badge}>Framework</p>
        <h1 className={styles.brand}>Starter Console</h1>

        <nav className={styles.nav}>
          <NavLink to="/" end className={navItemClassName}>Home</NavLink>
        </nav>
      </div>

      <div className={styles.logoutWrap}>
        <Button variant="secondary" onClick={handleLogout}>Logout</Button>
      </div>
    </aside>
  )
}
