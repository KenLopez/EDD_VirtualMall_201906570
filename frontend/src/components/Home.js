import React from 'react'
import { useHistory } from 'react-router'
import NavBar from './NavBar'
import StoreList from './StoreList'

function Home() {
    const history = useHistory()
    if (localStorage.getItem("LOGED") == null){
        history.push("/Login")
    }else if (localStorage.getItem("LOGED")=="Admin"){
        history.push("/Reporte")
    }
    return (
        <div>
            <NavBar
            activo={0}
            />
            <StoreList/>
        </div>
    )
}

export default Home
