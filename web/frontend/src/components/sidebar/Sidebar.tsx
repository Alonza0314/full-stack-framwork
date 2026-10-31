import { NavLink, useNavigate } from 'react-router-dom'
import Button from '../button/button'
import styles from './sidebar.module.css'

interface NavItem {
  to: string
  label: string
  // match only this exact path, not its children (e.g. "/")
  end?: boolean
}

interface NavGroup {
  // shown above the group's links; leave empty for the top group
  title?: string
  items: NavItem[]
}

// add pages here; a new group gets its own small heading in the sidebar
const navGroups: NavGroup[] = [
  {
    items: [
      { to: '/', label: 'Home', end: true },
    ],
  },
]

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

        {navGroups.map((group, index) => (
          <div key={group.title ?? index}>
            {group.title && <p className={styles.navGroup}>{group.title}</p>}
            <nav className={styles.nav}>
              {group.items.map((item) => (
                <NavLink key={item.to} to={item.to} end={item.end} className={navItemClassName}>{item.label}</NavLink>
              ))}
            </nav>
          </div>
        ))}
      </div>

      <div>
        <Button variant="secondary" onClick={handleLogout}>Logout</Button>
      </div>
    </aside>
  )
}
