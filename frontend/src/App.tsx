import { useState } from 'react'
import Button from './components/button/button'
import Modal from './components/modal/modal'
import Switch from './components/switch/switch'
import StatsCard from './components/stats/stats-card'
import NotificationContainer from './components/notifications/NotificationContainer'
import { useNotifications } from './hooks/useNotifications'
import './app.css'

export default function App() {
  const [modalOpen, setModalOpen] = useState(false)
  const [featureEnabled, setFeatureEnabled] = useState(false)
  const { errors, successes, addError, addSuccess, removeNotification } = useNotifications()

  return (
    <div className="app-shell">
      <NotificationContainer
        errors={errors}
        successes={successes}
        onClose={removeNotification}
      />

      <header className="app-header">
        <h1>Frontend Starter</h1>
        <p>Reusable components are ready. Build your new pages from here.</p>
      </header>

      <main className="app-main">
        <section className="demo-row">
          <StatsCard title="Build Status" value="Ready" description="Starter scaffold initialized" />
          <StatsCard title="Feature Flag" value={featureEnabled ? 'On' : 'Off'} description="Example state management" />
        </section>

        <section className="demo-actions">
          <Switch
            checked={featureEnabled}
            onChange={setFeatureEnabled}
            label="Enable demo feature"
          />
          <div className="button-row">
            <Button onClick={() => setModalOpen(true)}>Open Modal</Button>
            <Button variant="secondary" onClick={() => addSuccess('Saved successfully')}>
              Show Success
            </Button>
            <Button variant="secondary" onClick={() => addError('Something went wrong')}>
              Show Error
            </Button>
          </div>
        </section>
      </main>

      <Modal
        isOpen={modalOpen}
        onClose={() => setModalOpen(false)}
        title="Starter Modal"
        onSubmit={() => {
          addSuccess('Submitted from modal')
          setModalOpen(false)
        }}
      >
        <p>This modal component is reusable. Replace this content with your own forms or flows.</p>
      </Modal>
    </div>
  )
}
