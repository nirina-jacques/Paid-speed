import { Routes, Route } from 'react-router-dom'
import './App.css'
import Transaction from './transaction/Transaction'
import Login from './auth/login'
import Register from './auth/register'
import History from './history/History'

function App() {
  return (
    <Routes>
      <Route path="/" element={<Transaction />} />
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route path="/history" element={<History />} />
    </Routes>
  )
}

export default App
