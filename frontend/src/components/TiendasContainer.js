import React from 'react'
import { useHistory } from 'react-router'
import NavBar from './NavBar'
import StoreList from './StoreList'

function TiendasContainer() {
    const history = useHistory()
    if (localStorage.getItem("LOGED") == null){
        history.push("/Login")
    }else if (localStorage.getItem("LOGED")=="Cliente"){
        history.push("/Home")
    }
    return (
        <div>
            <NavBar
            activo={1}
            />
            <StoreList/>
        </div>
    )
}

export default TiendasContainer
