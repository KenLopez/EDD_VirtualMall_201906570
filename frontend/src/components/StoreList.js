import {React, useEffect, useState} from 'react'
import MosaicoTienda from './MosaicoTienda'
const axios = require('axios').default

function StoreList() {
    const [Datos, setDatos] = useState([])
    const [req, setreq] = useState(false)
    useEffect(() => {
        async function obtener(){
            if(!req){
                setreq(true)
                const data=await axios.get('http://localhost:3000/getTiendas')
                console.log(data.data.Datos)
                setDatos(data.data.Datos)
                console.log(Datos)
                console.log(req)
            }
        }
        obtener()
    }, [Datos, req])
    return (
        <div>
            <MosaicoTienda Datos={Datos}/>
        </div>
    )
}

export default StoreList
