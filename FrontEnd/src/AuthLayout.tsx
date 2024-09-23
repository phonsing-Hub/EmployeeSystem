
import { Outlet } from 'react-router-dom'

export default function AuthLayout() {
  return (
    <section className="flex justify-center items-center auth">
      <Outlet/>
      </section>
  )
}
