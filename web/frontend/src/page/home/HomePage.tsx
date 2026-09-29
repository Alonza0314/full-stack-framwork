import Sidebar from '../../components/sidebar/Sidebar'
import styles from './home-page.module.css'

export default function HomePage() {
  return (
    <div className={styles.layout}>
      <Sidebar />

      <main className={styles.content}>
        <header className={styles.header}>
          <h2>Home</h2>
          <p>A clean framework canvas ready for your features.</p>
        </header>

        <section className={styles.cardGrid}>
          <article className={styles.card}>
            <h3>Widget Area</h3>
            <p>Place dashboard cards, tables, or charts here.</p>
          </article>
          <article className={styles.card}>
            <h3>Feature Area</h3>
            <p>Use this section as a starting point for module pages.</p>
          </article>
        </section>
      </main>
    </div>
  )
}
