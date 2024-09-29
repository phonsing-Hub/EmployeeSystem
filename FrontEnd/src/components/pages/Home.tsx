import { useUser } from "../../UserProvider"
function Home() {
  const user = useUser()
  return (
    <div>Home{user?.email}</div>
  )
}

export default Home